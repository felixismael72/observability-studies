package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4/middleware"
	"go.elastic.co/apm/module/apmechov4"
)

func main() {
	app := LoadRoutes()
	app.Debug = true

	app.Use(middleware.Logger())

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	app.Use(apmechov4.Middleware())

	address := fmt.Sprintf("%s:%s", os.Getenv("SERVER_ADDRESS"), os.Getenv("SERVER_PORT"))

	app.Logger.Fatal(app.Start(address))
}
