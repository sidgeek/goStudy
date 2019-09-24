package main

import (
	"fmt"
	"log"
	"net/http"

	"goStudy/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	// connect db postgres
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=test password=example sslmode=disable")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	fmt.Println("sucess to connect database")
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&models.Product{})

	// Create
	db.Create(&models.Product{Code: "L1212", Price: 1000})

	// Read
	var product models.Product
	db.First(&product, 1) // find product with id 1
	fmt.Println("Product1", product)
	db.First(&product, "code = ?", "L1212") // find product with code l1212
	fmt.Println("Product2", product)

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	db.Delete(&product)

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":8080", router))
}
