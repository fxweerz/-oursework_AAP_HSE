package transport

import (
	"task-prioritizer/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", services.HomePage)
	return r
}
