package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/topnarapat/go-crud/configs"
	"github.com/topnarapat/go-crud/routes"
)

func main() {
	r := SetupRouter()

	r.Run(":" + os.Getenv("GO_PORT")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func SetupRouter() *gin.Engine {
	godotenv.Load(".env")

	configs.Connection()

	r := gin.Default()

	r.GET("/", routes.InitHomeRoutes)
	routes.InitUserRoutes(r)

	return r
}
