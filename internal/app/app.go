package app

import (
	"github.com/artemKapitonov/soundex/internal/handlers"
	"github.com/artemKapitonov/soundex/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type App struct {
	handler *handlers.Handler
	service *service.Service
	server  *gin.Engine
}

func New() (*App, error) {
	gin.SetMode(gin.DebugMode)

	logrus.SetFormatter(new(logrus.JSONFormatter))

	app := App{}

	app.service = service.New()

	app.handler = handlers.New(app.service)

	app.server = app.handler.InitRoutes()

	return &app, nil
}

func (a App) Run() error {

	if err := a.server.Run(":9000"); err != nil {
		logrus.Fatalf("can't start server: %s", err.Error())
	}

	return nil
}
