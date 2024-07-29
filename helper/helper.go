package helper

import "gorm.io/gorm"

const (
	Host     = "localhost"
	Port     = 5432
	User     = "postgres"
	Password = "password"
	Dbname   = "employee"
)

var DB *gorm.DB
