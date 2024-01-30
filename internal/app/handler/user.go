package handler

import (
	"ginconcurrency/internal/app/services"

	"net/http"
	"github.com/gin-gonic/gin"
)

type User interface {
	Get(context *gin.Context)
}

type user struct {
	userService services.User
}

func NewUser(userService services.User) User {
	return &user{userService: userService}
}

func (user *user)Get(context *gin.Context){
	users, err := user.userService.Get()
    if err!= nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    context.JSON(http.StatusOK, users)

}