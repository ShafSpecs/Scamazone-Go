package controllers

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
	"scamazone-Go/data/repositories"
	"scamazone-Go/services"
)

func CustomerController(r *mux.Router, db *gorm.DB) {
	customerRepo := repositories.NewCustomerRepository(db)
	_ = services.NewCustomerService(customerRepo)

	customerRouter := r.PathPrefix("/api/v1/customers").Subrouter()

	customerRouter.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		//_, _ = w.Write([]byte("Creating a customer"))
	}).Methods("GET")

	customerRouter.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		//_, _ = w.Write([]byte("Creating a customer"))
	}).Methods("GET")
}
