package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (a *TransactionHandler) Get2(c echo.Context) error {
	time.Sleep(1 * time.Second)

	message := "fuck"

	fmt.Println(message)

	return c.JSON(http.StatusOK, TransactionResponse{Message: message})
}
