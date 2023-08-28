package app

import (
	"avitoStart/internal/endpoint"
	"avitoStart/internal/service"
	"avitoStart/internal/storage/postgres"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
)

type App struct {
	database *postgres.Database
	endpoint *endpoint.Endpoint
	service  *service.Service
	echo     *echo.Echo
}

func New() (*App, error) {
	app := &App{}

	//dsn := os.Getenv("DSN")
	dsn := "host='127.0.0.1' port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"

	//инициализации всех параметров через New
	app.database, _ = postgres.New(dsn)
	app.service = service.New(app.database)
	app.endpoint = endpoint.New(app.service)
	app.echo = echo.New()

	//endpoints
	app.echo.POST("/user", app.endpoint.AddUser)
	app.echo.GET("/user", app.endpoint.ExtractUsers)
	app.echo.DELETE("/user", app.endpoint.DeleteUser)

	return app, nil
}
func (a *App) Run() error {
	fmt.Println("Server Runnig")

	err := a.echo.Start(":8090")
	if err != nil {
		log.Println(errors.New("Error Start Service"))
	}
	return nil
}
