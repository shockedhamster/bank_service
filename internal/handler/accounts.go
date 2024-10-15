package handler

import (
	"net/http"

	"github.com/bank_service/internal/entities"
	"github.com/gin-gonic/gin"
)

func (h *Handler) deposit(c *gin.Context) {
	var input entities.Operation

	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Account.Deposit(input.AccountId, input.Amount)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "deposit succesful",
	})
}

func (h *Handler) withdraw(c *gin.Context) {
	var input entities.Operation

	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Account.Withdraw(input.AccountId, input.Amount)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "withdraw succesful",
	})
}

func (h *Handler) transfer(c *gin.Context) {
	var input entities.Transfer

	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Account.Transfer(input.IdFrom, input.IdTo, input.Amount)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "transfer succesful",
	})
}
