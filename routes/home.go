package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitHomeRoutes(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"API VERSION": "1.0.0",
	})
}
