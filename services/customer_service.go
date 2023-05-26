package services

import (
	"scamazone-Go/data/models"
	"scamazone-Go/data/repositories"
)

type CustomerServiceImpl struct {
	customerRepo repositories.CustomerRepository
}

func NewCustomerService(customerRepo repositories.CustomerRepository) CustomerService {
	return &CustomerServiceImpl{customerRepo}
}

func (service *CustomerServiceImpl) Create(customer *models.Customer) error {
	return service.customerRepo.Create(customer)
}

func (service *CustomerServiceImpl) Update(customer models.Customer) error {
	return service.customerRepo.Update(customer)
}

func (service *CustomerServiceImpl) Delete(id uint) error {
	return service.customerRepo.Delete(id)
}

func (service *CustomerServiceImpl) FindById(id uint) (*models.Customer, error) {
	return service.customerRepo.FindById(id)
}

func (service *CustomerServiceImpl) FindByEmail(email string) (*models.Customer, error) {
	return service.customerRepo.FindByEmail(email)
}
