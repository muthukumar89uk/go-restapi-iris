package controllers

import (
	"fmt"
	"sampleIris/helper"
	"sampleIris/models"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

func CreateEmployee(ctx iris.Context) {
	var emp models.Employee

	if err := ctx.ReadJSON(&emp); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	uid := uuid.New()
	emp.Id = uid.String()

	if err := helper.DB.Debug().Create(&emp).Error; err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusOK)

	helper.DB.Debug().Preload("Addresses").Find(&emp, "id=?", emp.Id)
	ctx.JSON(emp)

}

func GetAllEmployee(ctx iris.Context) {
	var emp []models.Employee

	if err := helper.DB.Preload("Addresses").Find(&emp).Error; err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		fmt.Printf("failed to find employee: %v \n", err)
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(emp)
}

func GetEmpById(ctx iris.Context) {
	id := ctx.Params().Get("id")

	var emp models.Employee

	if err := helper.DB.Preload("Addresses").Find(&emp, "id=?", id).Error; err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		fmt.Printf("failed to find employee: %v \n", err)
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(emp)
}

func UpdateEmployee(ctx iris.Context) {
	empId := ctx.Params().Get("id")

	var input models.Employee

	if err := ctx.ReadJSON(&input); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	var employee models.Employee
	if err := helper.DB.Preload("Addresses").First(&employee, "id = ?", empId).Error; err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		fmt.Printf("failed to find employee: %v \n", err)
		return
	}

	// Update the fields
	if input.Name != "" {
		employee.Name = input.Name
	}
	if input.Addresses.City != "" {
		employee.Addresses.City = input.Addresses.City
	}
	if input.Addresses.State != "" {
		employee.Addresses.State = input.Addresses.State
	}
	if input.Addresses.Zip != 0 {
		employee.Addresses.Zip = input.Addresses.Zip
	}
	if input.Addresses.PhoneNumber != "" {
		employee.Addresses.PhoneNumber = input.Addresses.PhoneNumber
	}

	if err := helper.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(&employee).Error; err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		fmt.Printf("failed to update employee: %v \n", err)
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{
		"data":    employee,
		"message": "Employee updated successfully",
	})

}

func DeleteEmpById(ctx iris.Context) {
	id := ctx.Params().Get("id")

	var emp models.Employee

	if err := helper.DB.Preload("Addresses").Delete(&emp, "id=?", id).Error; err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		fmt.Printf("failed to find employee: %v \n", err)
		return
	}

	ctx.StatusCode(iris.StatusOK)

	ctx.JSON(iris.Map{
		"message": "data deleted successfully",
	})
}
