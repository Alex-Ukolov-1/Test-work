package main

import (
	"wb/app/wb/apiserver"
	"wb/app/wb/logger"
	"wb/app/wb/nats"
	"wb/app/wb/postgres"
)

func main() {
	go postgres.RecoverCash()
	go nats.Subscribe()
	if err := apiserver.Server.Start(); err != nil {
		logger.Log.Fatal(err)
	}
	return
}