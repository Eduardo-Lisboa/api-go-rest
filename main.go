package main

import (
	"github.com/Eduardo-Lisboa/go-res-api/database"
	"github.com/Eduardo-Lisboa/go-res-api/models"
	"github.com/Eduardo-Lisboa/go-res-api/routes"
)

func main() {

	models.Personalities = []models.Personality{
		{Id: 1, Name: "Eduardo", History: "I'm a software developer"},
		{Id: 2, Name: "Lisboa", History: "I'm a software developer"},
	}

	database.ConnectDatabase()
	routes.HandleRequest()
}
