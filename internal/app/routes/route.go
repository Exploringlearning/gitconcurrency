package routes

import (
	//"encoding/json"

	"ginconcurrency/internal/app/handler"
	"ginconcurrency/internal/app/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine)  {

    userService := services.NewUser()
	userHandler := handler.NewUser(userService)
	engine.GET("/", userHandler.Get)
}