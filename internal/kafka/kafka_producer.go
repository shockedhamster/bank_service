package kafka

import (
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

type KafkaProducer interface {
	SendMessage(topic string, message []byte) error
	Close() error
}

type Producer struct {
	producer sarama.SyncProducer
}

func NewKafkaProducer(cfg KafkaConfig) Producer {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{cfg.KafkaURL}, config)
	if err != nil {
		return Producer{}
	}
	return Producer{
		producer: producer,
	}
}

func (p *Producer) ProduceMessage(topic, key, value string) error {
	logrus.Infof("Data from ProduceMessage: USER_ID: %s, USERNAME: %s", key, value)

	message := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(value),
	}

	logrus.Infof("Data from ProduceMessage: USER_ID: %s, USERNAME: %s", key, value)

	partition, offset, err := p.producer.SendMessage(message)
	if err != nil {
		logrus.Printf("could not write message: %v", err)
		return err
	}

	logrus.Printf("message written successfully to partition %d at offset %d", partition, offset)
	return nil
}

func (p *Producer) Close() error {
	return p.producer.Close()
}
