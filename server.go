package main

import (
	v1 "github.com/anielski/download-url-in-go/models/v1"
	_ "math/big"
	"net/http"

	"github.com/anielski/download-url-in-go/app"
	"github.com/anielski/download-url-in-go/config"
	_ "github.com/anielski/download-url-in-go/docs"
	"github.com/anielski/download-url-in-go/handlers" // "fmt"

	"github.com/labstack/echo"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	config.LoadConfig()

	// Echo instance
	e := app.Init()

	h := &handlers.Handler{
		Gwp: new(v1.GWP),
	}

	//START WORKER
	v1.StartWorker( h.Gwp )

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	e.GET("/health-check", handlers.HealthCheck)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	api := e.Group("/api")

	api.POST("/fetcher", h.SaveFetcher)
	api.DELETE("/fetcher/:id", h.DeleteFetcher)
	api.GET("/fetcher/:id/history", h.GetHistory)
	api.GET("/fetcher", h.GetFetchers)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
