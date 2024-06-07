package main

import (
	"github.com/maxstefansson/boilerplate-api/api"
	_ "github.com/maxstefansson/boilerplate-api/docs"
)

// @title 		 Boilerplate API
// @version         1.0
// @description     This is a simple boilerplate API.
// @host      localhost:8080
// @BasePath  /api/v1

func main() {
	server := api.NewAPIServer(":8080")
	server.Run()
}
