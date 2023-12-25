package handler

import (
	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
}

type TransactionResponse struct {
	Message string `json:"message"`
	Body    string `json:"body,omitempty"`
}

func NewTransactionHandler(e *echo.Echo) {
	handler := &TransactionHandler{}

	e.GET("/get", handler.Get)
	e.GET("/get2", handler.Get2)
}
