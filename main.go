package main

import (
	"github.com/igormasi/api-go-gin/database"
	"github.com/igormasi/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
