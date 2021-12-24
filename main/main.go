package main

import (
	"log"
	"net/http"

	"github.com/Serenity0204/student-management-system-backend/packages/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	newRoute := mux.NewRouter()
	routes.RegisterStudentRoutes(newRoute)
	http.Handle("/", newRoute)
	log.Fatal(http.ListenAndServe("localhost:9000", newRoute))
}
