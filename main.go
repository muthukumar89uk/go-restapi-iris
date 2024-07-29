package main

import (
	"sampleIris/driver"
	"sampleIris/repository"
	"sampleIris/router"
)

func main() {
	driver.DbConnection()
	repository.Migration()
	router.Router()
}
