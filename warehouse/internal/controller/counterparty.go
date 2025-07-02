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

func CreateCounterparty(c *gin.Context) {
	// парсим в JSON
	var cp models.Counterparty
	if err := c.ShouldBindJSON(&cp); err != nil {
		HandleError(c, err)
		return
	}

	// проверяем валидации
	if cp.Name == "" ||
		cp.Contact == "" ||
		cp.Phone == "" ||
		cp.Email == "" {
		logger.Error.
			Printf("[controller] CreateCounterparty(): type in not provider/recipient or name, contact, phone or email is not filled: %s\n",
				errs.ErrBadRequestBody.Error(),
			)
		err := errors.Join(errors.New("type in not provider/recipient or name, contact, phone or email is not filled: "),
			errs.ErrBadRequestBody,
		)
		HandleError(c, err)
		return
	}

	// создаем
	cp, err := service.CreateCounterparty(cp)
	if err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusCreated, cp)
}

func GetCounterpartyByID(c *gin.Context) {
	// достаем id
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = errors.Join(err, errors.New("invalid id"), errs.ErrValidationFailed)
		HandleError(c, err)
		return
	}

	// достаем данные с бд
	cp, err := service.GetcounterpartyByID(id)
	if err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusOK, cp)
}

func GetAllCounterparties(c *gin.Context) {
	// достаем данные с бд
	cp, err := service.GetAllCounterparties()
	if err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusOK, cp)
}

func UpdateCounterpartyByID(c *gin.Context) {
	// достаем id
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = errors.Join(err, errors.New("invalid id"), errs.ErrValidationFailed)
		HandleError(c, err)
		return
	}

	// парсим в JSON
	var cp models.Counterparty
	if err = c.ShouldBindJSON(&cp); err != nil {
		HandleError(c, err)
		return
	}

	// пытаемся менять данные
	if err = service.UpdateCounterpartyByID(id, cp); err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusOK, gin.H{
		"message": "counterparty data updated successfully",
	})
}
