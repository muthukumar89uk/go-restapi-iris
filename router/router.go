package router

import (
	"sampleIris/controllers"

	"github.com/kataras/iris/v12"
)

func Router() {
	router := iris.New()

	router.Post("/v1/api/create/employee", controllers.CreateEmployee)
	router.Get("/v1/api/get/employees", controllers.GetAllEmployee)
	router.Get("/v1/api/getById/{id:uuid}", controllers.GetEmpById)
	router.Put("/v1/api/updateById/{id:uuid}", controllers.UpdateEmployee)
	router.Get("/v1/api/deleteById/{id:uuid}", controllers.DeleteEmpById)

	router.Run(iris.Addr(":2001"))
}
