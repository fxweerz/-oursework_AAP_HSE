package router

import (
	"task-prioritizer/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.LoadHTMLGlob("templates/go/*")
	r.Static("/static", "static")

	r.GET("/", handlers.Welcome)
	r.GET("/tasks", handlers.ListTasks)
	r.POST("/tasks/create", handlers.CreateTask)
	r.POST("/tasks/delete/:id", handlers.DeleteTask)

	return r
}
