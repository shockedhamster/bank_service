package kafka

import (
	"github.com/IBM/sarama"
	"github.com/sirupsen/logrus"
)

// func NewKafkaWriter(kafkaURL, topic string) *kafka.Writer {
// 	return kafka.NewWriter(kafka.WriterConfig{
// 		Brokers:  []string{kafkaURL},
// 		Topic:    topic,
// 		Balancer: &kafka.LeastBytes{},
// 	})
// }

// func ProduceMessage(writer *kafka.Writer, key, value string) error {
// 	logrus.Infof("Data from ProduceMessage: USER_ID: %s, USERNAME: %s", key, value)
// 	err := writer.WriteMessages(context.Background(),
// 		kafka.Message{
// 			Key:   []byte(key),
// 			Value: []byte(value),
// 		},
// 	)
// 	if err != nil {
// 		logrus.Printf("could not write message: %v", err)
// 		return err
// 	}
// 	logrus.Println("message written successfully")
// 	return nil
// }

func NewKafkaProducer(kafkaURL string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{kafkaURL}, config)
	if err != nil {
		return nil, err
	}
	return producer, nil
}

func ProduceMessage(producer sarama.SyncProducer, topic, key, value string) error {
	logrus.Infof("Data from ProduceMessage: USER_ID: %s, USERNAME: %s", key, value)

	message := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(value),
	}

	logrus.Infof("Data from ProduceMessage: USER_ID: %s, USERNAME: %s", key, value)

	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		logrus.Printf("could not write message: %v", err)
		return err
	}

	logrus.Printf("message written successfully to partition %d at offset %d", partition, offset)
	return nil
}
