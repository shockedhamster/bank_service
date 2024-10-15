package handler

import (
	"github.com/bank_service/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	operations := router.Group("/operations")
	{
		operations.GET("/user-balance/:id", h.getUserBalanceById)
		operations.GET("/transaction-history", h.getTransactionHistoryById)
	}

	account := router.Group("/account")
	{
		account.POST("/deposit", h.deposit)
		account.POST("/withdraw", h.withdraw)
		account.POST("/transfer", h.transfer)
	}

	return router
}
