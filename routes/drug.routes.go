package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jjmoreno-dev/technical-test-interfell/controllers"
	"github.com/jjmoreno-dev/technical-test-interfell/middleware"
)

type DrugRouteController struct {
	drugController controllers.Drug
}

func NewRouteDrugController(drugController controllers.Drug) DrugRouteController {
	return DrugRouteController{drugController}
}

func (drc *DrugRouteController) DrugRoute(rg *gin.RouterGroup) {

	router := rg.Group("drugs")
	router.Use(middleware.DeserializeUser())
	router.GET("/", drc.drugController.FindDrugs)
	router.POST("/", drc.drugController.CreateDrug)
	router.GET("/:drugId", drc.drugController.FindDrugById)
	router.DELETE("/:drugId", drc.drugController.DeleteDrug)
	router.PUT("/:drugId", drc.drugController.UpdateDrug)
}
