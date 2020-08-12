package main

import (
	"search/src/app"
	"search/src/config"
	"search/src/logger"
)

var c = config.GetConfig()
var log = logger.GetLogger("main")

func main() {
	log.Printf("search server is starting")
	a := app.NewApp(c)
	a.Run()
}
