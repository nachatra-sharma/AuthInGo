package main

import (
	"AuthInGo/app"
	config "AuthInGo/config/env"
)

func main() {

	appConfig := app.Config{
		Addr: config.GetString("PORT", ":8080"),
	}

	app := app.Application{
		Config: appConfig,
	}

	app.Run()
}
