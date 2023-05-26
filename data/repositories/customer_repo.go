package repositories

import (
	"gorm.io/gorm"
	"scamazone-Go/data/models"
)

type CustomerRepositoryImpl struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &CustomerRepositoryImpl{db}
}

func (repo *CustomerRepositoryImpl) Create(customer *models.Customer) error {
	return repo.db.Create(customer).Error
}

func (repo *CustomerRepositoryImpl) Update(customer models.Customer) error {
	return repo.db.Save(customer).Error
}

func (repo *CustomerRepositoryImpl) Delete(id uint) error {
	return repo.db.Delete(&models.Customer{}, id).Error
}

func (repo *CustomerRepositoryImpl) FindById(id uint) (*models.Customer, error) {
	var customer models.Customer
	err := repo.db.First(&customer, id).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (repo *CustomerRepositoryImpl) FindByEmail(email string) (*models.Customer, error) {
	var customer models.Customer
	err := repo.db.Where("email = ?", email).First(&customer).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
