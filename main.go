package main

import (
	"ip-data/internal/api"
	"ip-data/tools/wlog"
	"log"
)

func main() {
	wlog.New().Info("IP Data Server Starting")
	httpServer := api.NewServer()
	log.Fatal(httpServer.ListenAndServe())
}
