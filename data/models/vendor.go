package models

import "gorm.io/gorm"

type Vendor struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Password  string
}
