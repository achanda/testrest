package main

import (
	"github.com/achanda/testrest/app"
	"github.com/achanda/testrest/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
