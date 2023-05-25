package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jjmoreno-dev/technical-test-interfell/models"
	"gorm.io/gorm"
)

type Drug struct {
	DB *gorm.DB
}

func NewDrugController(DB *gorm.DB) Drug {
	return Drug{DB}
}

func (dc *Drug) FindDrugs(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var drugs []models.Drug
	results := dc.DB.Limit(intLimit).Offset(offset).Find(&drugs)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(drugs), "data": drugs})
}

func (dc *Drug) CreateDrug(ctx *gin.Context) {
	var payload *models.Drug

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newDrug := models.Drug{
		Name:        payload.Name,
		Approved:    payload.Approved,
		MinDose:     payload.MinDose,
		MaxDose:     payload.MaxDose,
		AvailableAt: payload.AvailableAt,
		CreatedAt:   time.Now(),
	}
	result := dc.DB.Create(&newDrug)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Drug with that title already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newDrug})
}

func (dc *Drug) FindDrugById(ctx *gin.Context) {
	drugId := ctx.Param("drugId")

	var drug models.Drug
	result := dc.DB.First(&drug, "id = ?", drugId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not drugs with that title exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": drug})
}

func (dc *Drug) DeleteDrug(ctx *gin.Context) {

	drugId := ctx.Param("drugId")
	var drug models.Drug

	existsDrug := dc.DB.First(&drug, "id = ?", drugId)
	if existsDrug.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not drug with that id exists"})
		return
	}

	result := dc.DB.Delete(&models.Drug{}, "id = ?", drugId)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Internal error occurred"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "the record was successfully deleted"})
}

func (dc *Drug) UpdateDrug(ctx *gin.Context) {
	drugId := ctx.Param("drugId")

	var payload *models.Drug
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updatedDrug models.Drug
	result := dc.DB.First(&updatedDrug, "id = ?", drugId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No drug with that title exists"})
		return
	}

	now := time.Now()
	drugToUpdate := models.Drug{
		Name:        payload.Name,
		Approved:    payload.Approved,
		MinDose:     payload.MinDose,
		MaxDose:     payload.MaxDose,
		AvailableAt: payload.AvailableAt,
		UpdatedAt:   &now,
	}
	dc.DB.Model(&updatedDrug).Updates(drugToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedDrug})
}
