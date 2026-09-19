package app

import (
	"AuthInGo/router"
	"fmt"
	"net/http"
	"time"
)


type Config struct {
	Addr string
}

type Application struct {
	Config Config
}

func (app *Application) Run() error {

	server := &http.Server{
		Addr: app.Config.Addr,
		Handler: router.SetupRouter(),
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Server is up and running on PORT", app.Config.Addr)

	return server.ListenAndServe()
}
