package main

import (
	"fitness-api/handlers"
	"fitness-api/storage"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	storage.InitDB()

	e.Use(handlers.LogRequest)

	e.GET("/", handlers.Home)
	e.POST("/users", handlers.CreateUser)
	e.POST("/measurements", handlers.CreateMeasurement)
	e.PUT("/users/:id", handlers.HandleUpdateUser)
	e.PUT("/measurements/:id", handlers.HandleUpdateMeasurement)

	e.Logger.Fatal(e.Start(":8080"))
}
