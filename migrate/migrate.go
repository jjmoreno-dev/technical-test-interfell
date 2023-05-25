package main

import (
	"fmt"
	"log"

	"github.com/jjmoreno-dev/technical-test-interfell/initializers"
	"github.com/jjmoreno-dev/technical-test-interfell/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{}, &models.Drug{}, &models.Vaccination{})
	fmt.Println("the migrations were executed satisfactorily")
}
