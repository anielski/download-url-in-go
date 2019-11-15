package handlers

import (
	"github.com/anielski/download-url-in-go/renderings"
	"net/http"
	"github.com/labstack/echo"
)

// HealthCheck - Health Check Handler
func HealthCheck(c echo.Context) error {
	resp := renderings.HealthCheckResponse{
		Message: "Everything is good!",
	}
	return c.JSON(http.StatusOK, resp)
}
