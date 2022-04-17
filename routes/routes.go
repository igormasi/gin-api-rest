package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/igormasi/api-go-gin/controllers"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.Run()
}
