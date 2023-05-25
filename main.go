package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jjmoreno-dev/technical-test-interfell/controllers"
	"github.com/jjmoreno-dev/technical-test-interfell/initializers"
	"github.com/jjmoreno-dev/technical-test-interfell/routes"
)

var (
	server              *gin.Engine
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	DrugController      controllers.Drug
	DrugRouteController routes.DrugRouteController

	VaccinationController      controllers.Vaccination
	VaccinationRouteController routes.VaccinationRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	DrugController = controllers.NewDrugController(initializers.DB)
	DrugRouteController = routes.NewRouteDrugController(DrugController)

	VaccinationController = controllers.NewVaccinationController(initializers.DB)
	VaccinationRouteController = routes.NewRouteVaccinationController(VaccinationController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")

	AuthRouteController.AuthRoute(router)
	DrugRouteController.DrugRoute(router)
	VaccinationRouteController.VaccinationRoute(router)

	log.Fatal(server.Run(":" + config.ServerPort))
}
