package services

import (
	"scamazone-Go/data/repositories"
)

type CustomerServiceImpl struct {
	customerRepo repositories.CustomerRepository
}

func NewCustomerService(customerRepo repositories.CustomerRepository) CustomerService {
	return &CustomerServiceImpl{customerRepo}
}
