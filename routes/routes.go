package routes

import (
	"api-go-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandlerRequest() {
	r := gin.Default()
	r.GET("/alunos", controllers.ShowAllStudent)
	r.GET("/:name", controllers.Gretting)
	r.POST("/student", controllers.CreateStudent)
	r.GET("/student", controllers.ListStundent)
	r.GET("/student/:id", controllers.GetStudent)
	r.DELETE("/student/:id", controllers.DeleteStudent)
	r.Run()
}
