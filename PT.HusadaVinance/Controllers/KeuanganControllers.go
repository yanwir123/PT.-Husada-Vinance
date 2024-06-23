package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	models "PT.HusadaVinance/Models"
	DataPerusahaan "PT.HusadaVinance/Models/DataPerusahaan"
	KeuanganRepository "PT.HusadaVinance/repository"
)

// InsertJurusanController adalah handler untuk menyisipkan data keuangan baru.
func GetHusadaVinanceControllersByID(c *gin.Context) {
	var request DataPerusahaan.Keuangan
	var response models.BaseResponseModels

	response = KeuanganRepository.GetHusadaVinanceByID(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func InsertHusadaVinanceControllers(c *gin.Context) {
	var request DataPerusahaan.Keuangan
	var response models.BaseResponseModels

	if err := c.ShouldBindJSON(&request); err != nil {
		response = models.BaseResponseModels{
			CodeResponse:  400,
			HeaderMessage: "Bad Request",
			Message:       err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response = KeuanganRepository.InsertHusadaVinance(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func UpdateHusadavinanceControllers(c *gin.Context) {
	var request DataPerusahaan.Keuangan
	var response models.BaseResponseModels

	if err := c.ShouldBindJSON(&request); err != nil {
		response = models.BaseResponseModels{
			CodeResponse:  400,
			HeaderMessage: "Bad Request",
			Message:       err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response = KeuanganRepository.UpdateHusadainance(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func DeleteHusadaVinanceControllers(c *gin.Context) {
	var request DataPerusahaan.Keuangan
	var response models.BaseResponseModels

	if err := c.ShouldBindJSON(&request); err != nil {
		response = models.BaseResponseModels{
			CodeResponse:  400,
			HeaderMessage: "Bad Request",
			Message:       err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response = KeuanganRepository.DeleteHusadaVinance(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}