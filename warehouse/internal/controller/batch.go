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

func CreateBatch(c *gin.Context) {
	// парсинг из json
	var b models.Batch
	if err := c.ShouldBindJSON(&b); err != nil {
		HandleError(c, err)
		return
	}

	// проверка валидаций
	if (b.Type != "in" && b.Type != "out") ||
		!b.Date.IsZero() ||
		b.CounterpartyName == "" ||
		b.Article == "" ||
		b.Quantity <= 0 {
		logger.Error.
			Printf("[controller] CreateBatch(): one or more fields are filled in incorrectly: %s\n", errs.ErrBadRequestBody.Error())
		err := errors.Join(errors.New("one or more fields are filled in incorrectly"), errs.ErrBadRequestBody)
		HandleError(c, err)
		return
	}
	b.Username = c.GetString(userUsernameCtx)

	// попытка протолкнуть в сервис
	err := service.CreateBatch(b)
	if err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusCreated, gin.H{
		"message": "batch created successfully",
	})
}

func GetAllBatches(c *gin.Context) {
	// достаем данные с бд
	batches, err := service.GetAllBatches()
	if err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusOK, batches)
}

func GetBatchByID(c *gin.Context) {
	// доствем if
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		err = errors.Join(err, errors.New("invalid id"), errs.ErrValidationFailed)
		HandleError(c, err)
		return
	}

	// достаем данные с бд
	batch, err := service.GetBatchByID(id)
	if err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusOK, batch)
}
