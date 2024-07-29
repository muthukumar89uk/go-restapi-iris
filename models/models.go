package models

type Employee struct {
	Id        string  `gorm:"column:id;omitempty"`
	Name      string  `json:"name" gorm:"column:name;omitempty"`
	Addresses Address `json:"address" gorm:"foreignKey:EmpId;omitempty"`
}

type Address struct {
	Id          uint   `gorm:"column:id;omitempty"`
	EmpId       string `json:"-" gorm:"omitempty"`
	City        string `json:"city" gorm:"column:city;omitempty"`
	State       string `json:"state" gorm:"column:state;omitempty"`
	Zip         int    `json:"zip" gorm:"column:zip;omitempty"`
	PhoneNumber string `json:"phone_No" gorm:"column:phone_No;omitempty"`
}

func (Employee) TableName() string {
	return "employee"
}

func (Address) TableName() string {
	return "address"
}
