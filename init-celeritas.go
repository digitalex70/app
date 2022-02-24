package main

import (
	"log"
	"os"

	"app/handlers"

	"github.com/digitalex70/celeritas"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	cel := &celeritas.Celeritas{}
	err = cel.New(path)
	if err != nil {
		log.Fatal(err)
	}

	cel.AppName = "app"

	myHandlers := &handlers.Handlers{
		App: cel,
	}

	cel.InfoLog.Println("Debug is set to:", cel.Debug)

	app := &application{
		App:      cel,
		Handlers: myHandlers,
	}
	app.App.Routes = app.routes()
	return app
}
