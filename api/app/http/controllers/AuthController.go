package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type AuthController struct {

}

type UserLogin struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (*AuthController) Login(ctx *gin.Context) {
	var userLogin UserLogin
	if err := ctx.ShouldBindJSON(&userLogin); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(userLogin)
}

func NewAuthController() *AuthController {
	return &AuthController{}
}