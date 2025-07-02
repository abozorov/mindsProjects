package controller

import (
	"net/http"
	"warehouse/internal/service"

	"github.com/gin-gonic/gin"
)

func GetStorageByAdressCode(c *gin.Context) {
	// достаем adressCode
	adressCode := c.Param("adressCode")

	// достаем данные с бд
	storage, err := service.GetStorageByAdressCode(adressCode)
	if err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusOK, storage)

}

func GetAllStorages(c *gin.Context) {
	// достаем данные с бд
	storages, err := service.GetAllStorages()
	if err != nil {
		HandleError(c, err)
		return
	}

	// ответ
	c.JSON(http.StatusOK, storages)
}
