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

func CreateProduct(c *gin.Context) {
	// парсим в JSON
	var p models.Product

	if err := c.ShouldBindJSON(&p); err != nil {
		HandleError(c, err)
		return
	}

	// проверяем валидации
	if p.Article == "" {
		logger.Error.
			Printf("[controller] CreateProduct(): article is not filled in: %s\n", errs.ErrBadRequestBody.Error())
		err := errors.Join(errors.New("article is not filled in: "), errs.ErrBadRequestBody)
		HandleError(c, err)
		return
	}

	// создаем
	p, err := service.CreateProduct(p)
	if err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusCreated, p)

}

func GetProductByID(c *gin.Context) {
	// достаем id
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = errors.Join(err, errors.New("invalid id"), errs.ErrValidationFailed)
		HandleError(c, err)
		return
	}

	// достаем данные с бд
	product, err := service.GetProductByID(id)
	if err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusOK, product)

}

func GetAllProducts(c *gin.Context) {
	// достаем данные с бд
	products, err := service.GetAllProducts()
	if err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusOK, products)
}

func UpdateProductByID(c *gin.Context) {
	// достаем id
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = errors.Join(err, errors.New("invalid id"), errs.ErrValidationFailed)
		HandleError(c, err)
		return
	}

	// парсим в JSON
	var p models.Product

	if err := c.ShouldBindJSON(&p); err != nil {
		HandleError(c, err)
		return
	}

	// пытаемся менять данные
	if err = service.UpdateProductByID(id, p); err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusOK, gin.H{
		"message": "product data updated successfully",
	})
}
