package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jjmoreno-dev/technical-test-interfell/controllers"
	"github.com/jjmoreno-dev/technical-test-interfell/middleware"
)

type VaccinationRouteController struct {
	vaccinationController controllers.Vaccination
}

func NewRouteVaccinationController(vaccinationController controllers.Vaccination) VaccinationRouteController {
	return VaccinationRouteController{vaccinationController}
}

func (drc *VaccinationRouteController) VaccinationRoute(rg *gin.RouterGroup) {

	router := rg.Group("vaccinations")
	router.Use(middleware.DeserializeUser())
	router.GET("/", drc.vaccinationController.FindVaccinations)
	router.POST("/", drc.vaccinationController.CreateVaccination)
	router.GET("/:vaccinationId", drc.vaccinationController.FindVaccinationById)
	router.DELETE("/:vaccinationId", drc.vaccinationController.DeleteVaccination)
	router.PUT("/:vaccinationId", drc.vaccinationController.UpdateVaccination)
}
