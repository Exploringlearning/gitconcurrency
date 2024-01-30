package main

import (
	//"fmt"
	"ginconcurrency/internal/app/routes"

	"github.com/gin-gonic/gin"
)

	  func main() {
		
		engine := gin.Default()
		routes.RegisterRoutes(engine)
		engine.Run(":8080")
	  }
	