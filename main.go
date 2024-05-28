package main

import (
	"github.com/MarioGalvaoWoohoo/api-go-gin/database"
	"github.com/MarioGalvaoWoohoo/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
