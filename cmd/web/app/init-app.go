package app

import (
	"baboonapp/database/models"
	"baboonapp/http/handlers"
	"baboonapp/http/middleware"
	"log"
	"os"

	"github.com/martijnkorbee/gobaboon"
	"github.com/martijnkorbee/gobaboon/pkg/logger"
)

func MustInitApplication() *application {
	// get application root path
	path, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	// init baboon
	baboon := &gobaboon.Baboon{}
	err = baboon.Init(path)
	if err != nil {
		log.Fatalln(err)
	}

	app := &application{
		Baboon: baboon,
	}

	// start application logger
	log := &logger.LoggerConfig{
		Rootpath:   app.Baboon.Config.Rootpath,
		Debug:      app.Baboon.Config.Debug,
		Console:    app.Baboon.Config.Debug,
		ToFile:     !app.Baboon.Config.Debug,
		Service:    "application-main",
		Filename:   "/logs/app_log.log",
		MaxBackups: 2,
		LocalTime:  true,
	}
	app.Log = log.Start()

	// add models
	app.Models = models.New(app.Baboon.Database)

	// add middleware
	app.Middleware = &middleware.Middleware{
		App:    baboon,
		Models: app.Models,
	}

	// add handlers
	app.Handlers = &handlers.Handlers{
		App: baboon,
	}

	// mount application routes
	app.Baboon.Server.Router.Mount("/", app.routes())       // web routes
	app.Baboon.Server.Router.Mount("/api", app.routesAPI()) // API routes

	return app
}
