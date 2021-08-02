package main

import (
	"github.com/adrianpanatra/adrian-bank/app"
	"github.com/adrianpanatra/adrian-bank/logger"
)

func main() {

	// log.Println("starting our application...")
	logger.Info("Starting the application..")
	app.Start()
}
