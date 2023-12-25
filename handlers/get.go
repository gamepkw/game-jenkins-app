package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

func (a *TransactionHandler) Get(c echo.Context) error {
	time.Sleep(1 * time.Second)

	if err := callServiceB(); err != nil {
		return c.JSON(http.StatusInternalServerError, TransactionResponse{Message: err.Error()})
	}

	message := "Call service B success fuck"

	fmt.Println(message)

	return c.JSON(http.StatusOK, TransactionResponse{Message: message})
}

func callServiceB() error {
	httpClient := &http.Client{}
	// endpoint := "http://localhost:8083/get"
	serviceName := "test-service-b-cluster-ip-service"
	namespace := "test-service-b-namespace"
	fmt.Println(namespace, serviceName)
	endpoint := fmt.Sprintf("http://%s.%s.svc.cluster.local:80/get", serviceName, namespace)
	fmt.Println(endpoint)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return errors.Wrap(err, "error make http request")
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "error get http response")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to call service b, status code: %d", resp.StatusCode)
	}

	fmt.Println(resp.StatusCode)

	return nil
}
