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

type Vaccination struct {
	DB *gorm.DB
}

func NewVaccinationController(DB *gorm.DB) Vaccination {
	return Vaccination{DB}
}

func (dc *Vaccination) FindVaccinations(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var vaccinations []models.Vaccination
	results := dc.DB.Limit(intLimit).Offset(offset).Find(&vaccinations)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(vaccinations), "data": vaccinations})
}

func (vc *Vaccination) CreateVaccination(ctx *gin.Context) {
	var payload *models.Vaccination

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	newVaccination := models.Vaccination{
		Name:      payload.Name,
		Dose:      payload.Dose,
		DrugId:    payload.DrugId,
		Date:      payload.Date,
		CreatedAt: time.Now(),
	}
	result := vc.DB.Create(&newVaccination)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Vaccination with that title already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newVaccination})
}

func (dc *Vaccination) FindVaccinationById(ctx *gin.Context) {
	vaccinationId := ctx.Param("vaccinationId")

	var vaccination models.Vaccination
	result := dc.DB.First(&vaccination, "id = ?", vaccinationId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not vaccination with that title exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": vaccination})
}

func (dc *Vaccination) DeleteVaccination(ctx *gin.Context) {

	vaccinationId := ctx.Param("drugId")
	var vaccination models.Vaccination

	existsVaccination := dc.DB.First(&vaccination, "id = ?", vaccinationId)
	if existsVaccination.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Not Vaccination with that id exists"})
		return
	}

	result := dc.DB.Delete(&models.Vaccination{}, "id = ?", vaccinationId)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Internal error occurred"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "the record was successfully deleted"})
}

func (dc *Vaccination) UpdateVaccination(ctx *gin.Context) {
	vaccinationId := ctx.Param("vaccinationId")

	var payload *models.Vaccination
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updatedVaccination models.Vaccination
	result := dc.DB.First(&updatedVaccination, "id = ?", vaccinationId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No Vaccination with that title exists"})
		return
	}

	now := time.Now()
	drugToUpdate := models.Drug{
		Name: payload.Name,
		/*Approved:    payload.Approved,
		MinDose:     payload.MinDose,
		MaxDose:     payload.MaxDose,
		AvailableAt: payload.AvailableAt,
		*/
		UpdatedAt: &now,
	}
	dc.DB.Model(&updatedVaccination).Updates(drugToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedVaccination})
}
