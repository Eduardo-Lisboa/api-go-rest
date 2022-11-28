package main

import (
	"fmt"

	"github.com/Eduardo-Lisboa/go-res-api/database"
	"github.com/Eduardo-Lisboa/go-res-api/routes"
)

func main() {

	database.ConnectDatabase()
	fmt.Println("Database connected")
	routes.HandleRequest()
}
