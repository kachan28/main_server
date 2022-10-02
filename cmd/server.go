package cmd

import (
	"log"
	"main_server/app"
	"main_server/config"
)

func Execute() {
	conf, err := config.InitializeConfig()
	if err != nil {
		panic(err)
	}

	server, err := app.Init(conf)
	if err != nil {
		panic(err)
	}

	log.Println(server.Start(conf))
}
