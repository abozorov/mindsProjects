package controller

import (
	"errors"
	"net/http"
	"strconv"
	"warehouse/internal/errs"
	"warehouse/internal/models"
	"warehouse/internal/service"
	"warehouse/logger"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	// парсим в JSON
	var u models.User

	if err := c.ShouldBindJSON(&u); err != nil {
		HandleError(c, err)
		return
	}

	// проверяем валидации
	if u.Username == "" || u.Password == "" || u.FullName == "" {
		logger.Error.
			Printf("[controller] CreateUser(): username, password or full name is not filled in: %s\n",
				errs.ErrBadRequestBody.Error(),
			)
		err := errors.Join(errors.New("username, password or full name is not filled in: "),
			errs.ErrBadRequestBody,
		)
		HandleError(c, err)
		return
	}

	// создаем юзера
	user, err := service.CreateUser(u)
	if err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusCreated, user)
}

func GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = errors.Join(err, errors.New("invalid id"), errs.ErrValidationFailed)
		HandleError(c, err)
		return
	}

	user, err := service.GetUserByID(id)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetAllUsers(c *gin.Context) {
	users, err := service.GetAllUsers()
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func UpdateUserByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = errors.Join(err, errors.New("invalid id"), errs.ErrValidationFailed)
		HandleError(c, err)
		return
	}

	var u models.User
	if err = c.ShouldBindJSON(&u); err != nil {
		HandleError(c, err)
		return
	}

	if err = service.UpdateUserByID(id, u); err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "user data updated successfully",
	})
}
