package main

import (
	"github.com/igormasi/api-go-gin/database"
	"github.com/igormasi/api-go-gin/models"
	"github.com/igormasi/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "Igor Silva", CPF: "00000000000", RG: "4700000000"},
		{Nome: "Ana Silva", CPF: "11111111111", RG: "4800000000"},
	}
	routes.HandleRequest()
}
