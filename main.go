package main

import (
	"github.com/gin-gonic/gin"
	"user-service/config/boot"
	cfg "user-service/config/env"
	"user-service/routes"
)

type App struct {
	config cfg.Config
}

var	(
	app App
)

func init() {
	config := cfg.NewViperConfig()
	app = App{config: config}
}

func main() {
	e := gin.Default()

	handler := boot.ConfigHandler{
		E: e,
		Config: app.config,
	}
	useCase := handler.RegisterConfig()

	routes.ServiceRoute(handler, useCase)
}