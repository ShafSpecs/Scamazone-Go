package services

import (
	"scamazone-Go/dtos/requests"
	"scamazone-Go/dtos/responses"
)

type CustomerService interface {
	CreateCustomer(request requests.CustomerRegisterRequest) (responses.CustomerRegisterResponse, error)
}
