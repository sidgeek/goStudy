// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"goStudy/models"

// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/postgres"
// 	"github.com/julienschmidt/httprouter"
// )

// func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
// 	fmt.Fprint(w, "Welcome!\n")
// }

// func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
// }

// func main() {
// 	// connect db postgres
// 	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=test password=example sslmode=disable")
// 	if err != nil {
// 		fmt.Println(err)
// 		panic("failed to connect database")
// 	}
// 	fmt.Println("sucess to connect database")
// 	defer db.Close()

// 	user := models.User{LoginName: "Jinzhu", Pwd: "123456"}
// 	db.NewRecord(user) // => returns `true` as primary key is blank
// 	db.Create(&user)

// 	router := httprouter.New()
// 	router.GET("/", Index)
// 	router.GET("/hello/:name", Hello)

// 	log.Fatal(http.ListenAndServe(":8080", router))
// }
