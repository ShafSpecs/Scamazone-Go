package services

import "scamazone-Go/data/models"

// CustomerService change this to normal funcs like login, etc.
type CustomerService interface {
	Create(customer *models.Customer) error
	Update(customer models.Customer) error
	Delete(id uint) error
	FindById(id uint) (*models.Customer, error)
	FindByEmail(email string) (*models.Customer, error)
}

type VendorService interface {
	Create(vendor *models.Vendor) error
	Update(vendor models.Vendor) error
	Delete(id uint) error
	FindById(id uint) (*models.Vendor, error)
	FindByEmail(email string) (*models.Vendor, error)
}
