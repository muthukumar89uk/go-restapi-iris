package router

import (
	"sampleIris/controllers"

	"github.com/kataras/iris/v12"
)

func Router() {
	router := iris.New()

	router.Post("/create", controllers.CreateEmployee)
	router.Get("/getAllEmp", controllers.GetAllEmployee)
	router.Get("/getById/{id:uuid}", controllers.GetEmpById)
	router.Put("updateById/{id:uuid}", controllers.UpdateEmployee)
	router.Get("/delete/{id:uuid}", controllers.DeleteEmpById)

	router.Run(iris.Addr(":2001"))
}
