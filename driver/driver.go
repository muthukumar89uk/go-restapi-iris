package driver

import (
	"fmt"
	"log"
	"sampleIris/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error

func DbConnection() {
	connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", helper.Host, helper.Port, helper.User, helper.Password, helper.Dbname)
	helper.DB, err = gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		log.Println("Db Connection is failed", err)
	}
	fmt.Printf("%s Database connected \n", helper.Dbname)

}
