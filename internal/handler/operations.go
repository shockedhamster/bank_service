package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bank_service/internal/entities"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getUserBalanceById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "invalid id param")
		fmt.Println(id)
		return
	}
	balance, err := h.services.Operations.GetUserBalanceById(id)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"balance": balance,
	})
}

// ? Или сделать через параметры в запросе get
func (h *Handler) getTransactionHistoryById(c *gin.Context) {
	var input entities.GetTransactionHistoryInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	resultTransactionList, err := h.services.Operations.GetTransactionHistoryById(input.Id, input.SortType, input.Limit, input.Offset)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	type responce struct {
		TransactionList []entities.Operation `json:"transaction list"`
	}

	c.JSON(http.StatusOK, responce{
		TransactionList: resultTransactionList,
	})

}
