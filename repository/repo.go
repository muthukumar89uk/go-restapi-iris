package repository

import (
	"fmt"
	"sampleIris/helper"
	"sampleIris/models"
)

func Migration() {
	err := helper.DB.AutoMigrate(&models.Employee{}, &models.Address{})
	if err != nil {
		fmt.Println("Failed to create Tables")
		return
	}
}
