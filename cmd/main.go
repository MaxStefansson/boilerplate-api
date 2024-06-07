package main

import (
	"github.com/maxstefansson/boilerplate-api/api"
)

func main() {
	server := api.NewAPIServer(":8080")
	server.Run()
}
