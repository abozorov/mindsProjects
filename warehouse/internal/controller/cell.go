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

func CreateCell(c *gin.Context) {
	// парсим в JSON
	var cell models.Cell

	if err := c.ShouldBindJSON(&cell); err != nil {
		HandleError(c, err)
		return
	}

	// проверяем валидации
	if cell.Zone == "" ||
		cell.Row == 0 ||
		cell.AdressCode == "" {
		logger.Error.
			Printf("[controller] CreateCell(): zone, row or adress code is not filled in: %s\n", errs.ErrBadRequestBody.Error())
		err := errors.Join(errors.New("zone, row or adress code is not filled in: "), errs.ErrBadRequestBody)
		HandleError(c, err)
		return
	}

	// создаем
	cell, err := service.CreateCell(cell)
	if err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusCreated, cell)

}

func GetCellByID(c *gin.Context) {
	// достаем id
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = errors.Join(err, errors.New("invalid id"), errs.ErrValidationFailed)
		HandleError(c, err)
		return
	}

	// достаем данные с бд
	Cell, err := service.GetCellByID(id)
	if err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusOK, Cell)

}

func GetAllCells(c *gin.Context) {
	// достаем данные с бд
	cells, err := service.GetAllCells()
	if err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusOK, cells)
}

func UpdateCellByID(c *gin.Context) {
	// достаем id
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = errors.Join(err, errors.New("invalid id"), errs.ErrValidationFailed)
		HandleError(c, err)
		return
	}

	// парсим в JSON
	var cell models.Cell

	if err := c.ShouldBindJSON(&cell); err != nil {
		HandleError(c, err)
		return
	}

	// пытаемся менять данные
	if err = service.UpdateCellByID(id, cell); err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusOK, gin.H{
		"message": "cell data updated successfully",
	})
}
