package kafka_consumer

import (
	"context"
	"strconv"

	"github.com/IBM/sarama"
	"github.com/bank_service/internal/service"
	"github.com/sirupsen/logrus"
)

func ConsumeMessages(kafkaURL string, topic string, accountService *service.Service) {
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()

	consumerGroup, err := sarama.NewConsumerGroup([]string{kafkaURL}, "bank_service_group", config)
	if err != nil {
		logrus.Fatalf("Error creating consumer group: %v", err)
	}
	defer consumerGroup.Close()

	ctx := context.Background()

	for {
		err := consumerGroup.Consume(ctx, []string{topic}, &consumerGroupHandler{accountService: accountService})
		if err != nil {
			logrus.Errorf("Error from consumer: %v", err)
		}
	}
}

// consumerGroupHandler обрабатывает сообщения из Consumer Group
type consumerGroupHandler struct {
	accountService *service.Service
}

// Setup выполняется перед запуском Consumer Group
func (h *consumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup выполняется после завершения работы Consumer Group
func (h *consumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim обрабатывает сообщения из Consumer Group
func (h *consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		logrus.Printf("received message: key=%s, value=%s", string(message.Key), string(message.Value))

		id, err := strconv.Atoi(string(message.Key))
		if err != nil {
			logrus.Errorf("Error converting msg to int: %s", err.Error())
			continue
		}

		err = h.accountService.CreateAccount(id)
		if err != nil {
			logrus.Printf("Error while creating account: %v", err)
		}

		session.MarkMessage(message, "")
	}
	return nil
}
