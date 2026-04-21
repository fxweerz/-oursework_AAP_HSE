package router

import (
	"task-prioritizer/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

	r.LoadHTMLGlob("static/*")

	r.GET("/", handlers.ListTasks)
	r.POST("/task", handlers.CreateTask)

	return r
}
