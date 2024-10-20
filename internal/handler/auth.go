package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/bank_service/internal/entities"
	"github.com/bank_service/internal/kafka"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) signUp(c *gin.Context) {
	var input entities.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Создаем Kafka writer
	producer, err := kafka.NewKafkaProducer("kafka:9092")
	if err != nil {
		logrus.Errorf("Failed to init producer: %s", err.Error())
	}
	defer producer.Close()

	idStr := strconv.Itoa(id)

	logrus.Infof("Message is ready to send to Kafka (in converted): USERID: %s, USERNAME:%s", idStr, input.Username)

	// Отправляем сообщение в Kafka
	topic := "user-created"
	if err := kafka.ProduceMessage(producer, topic, idStr, input.Username); err != nil {
		log.Printf("Error while producing message to kafka: %v", err)
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})

}

func (h *Handler) signIn(c *gin.Context) {
	var input entities.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
