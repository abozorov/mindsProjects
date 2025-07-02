package controller

import (
	"net/http"
	"warehouse/internal/models"
	"warehouse/internal/service"
	"warehouse/utils"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	// получить идентификатор и пароль
	var u models.UserSignIn
	if err := c.ShouldBindJSON(&u); err != nil {
		HandleError(c, err)
		return
	}

	// отправить в бд запрос на проверку есть ли такой
	// пользователь с таким паролем
	user, err := service.GetUserByUsernameAndPassword(u.Username, u.Password)
	if err != nil {
		HandleError(c, err)
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": token,
	})
}
