package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// начальная страница / main page
func HomePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "67"})
}
