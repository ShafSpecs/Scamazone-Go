package main

import (
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"scamazone-Go/controllers"
)

func main() {
	r := mux.NewRouter()

	dsn := "root:Abdur-Rahman05@tcp(127.0.0.1:3306)/scamazone?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	//r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	_, _ = fmt.Fprintf(w, "Hello from: '%s' route", r.URL.String())
	//})
	//
	//r.HandleFunc("/{title}/{page}", func(w http.ResponseWriter, r *http.Request) {
	//	vars := mux.Vars(r)
	//	title := vars["title"]
	//	page := vars["page"]
	//
	//	_, _ = fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	//})

	controllers.CustomerController(r, db)

	log.Fatal(http.ListenAndServe(":8000", r))
}
