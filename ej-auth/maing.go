package main

import (
	"github.com/AgosHenderson/ArqSoftware1/ej-auth/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	engine.POST("/login",
		controllers.Login)
	engine.Run(":8080")
}
