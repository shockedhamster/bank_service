package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/bank_service/internal/entities"
	"github.com/gin-gonic/gin"
)

const (
	getCurrencyDataUrl = "https://api.currencyapi.com/v3/latest"
	currencyRUB        = "RUB"
	currencyUSD        = "USD"
	currencyEUR        = "EUR"
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

	currency := c.Param("currency")
	if currency != currencyRUB {
		// Вынести загрузку апи в отдельный сервис
		currencyApiKey := os.Getenv("CURRENCY_API_KEY")

		requestData := entities.CurrencyReq{
			Currencies:    currency,
			Base_currency: currencyRUB,
		}

		reqBody, _ := json.Marshal(requestData)

		req, err := http.NewRequest("GET", getCurrencyDataUrl, bytes.NewBuffer(reqBody))
		if err != nil {
			newErrorResponce(c, http.StatusInternalServerError, err.Error())
			return
		}

		req.Header.Set("apikey", currencyApiKey)
		req.Header.Set("Content-Type", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			newErrorResponce(c, http.StatusInternalServerError, err.Error())
			return
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			newErrorResponce(c, http.StatusInternalServerError, err.Error())
			return
		}

		var currenciesResponse entities.CurrencyResp
		err = json.Unmarshal(body, &currenciesResponse)
		if err != nil {
			newErrorResponce(c, http.StatusInternalServerError, err.Error())
			return
		}

		// Надо бы сделать отдельную сущность на каждый ответ
		switch currency {
		case currencyUSD:
			c.JSON(http.StatusOK, gin.H{
				"currency": currenciesResponse.Data.USD.Code,
				"balance":  float64(balance) * currenciesResponse.Data.USD.Value,
			})
		case currencyEUR:
			c.JSON(http.StatusOK, gin.H{
				"currency": currenciesResponse.Data.EUR.Code,
				"balance":  float64(balance) * currenciesResponse.Data.EUR.Value,
			})
		default:
			c.JSON(http.StatusOK, gin.H{
				"message": "unknown currency",
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"currency": currencyRUB,
			"balance":  balance,
		})
	}
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
