package main

import (
	"sample-server/internal/app"
	"sample-server/internal/config"
)

func main() {
	conf := config.LoadFromEnv()
	e := app.NewServer(conf)
	e.Logger.Fatal(e.Start(":" + conf.Port))
}
