package routes

import (
	"github.com/Serenity0204/student-management-system-backend/packages/controllers"
	"github.com/gorilla/mux"
)

var RegisterStudentRoutes = func(router *mux.Router) {
	router.HandleFunc("/student/", controllers.CreateStudent).Methods("POST")
	router.HandleFunc("/student/", controllers.GetStudent).Methods("GET")
	router.HandleFunc("/student/{studentId}", controllers.GetStudentById).Methods("GET")
	router.HandleFunc("/student/{studentId}", controllers.UpdateStudent).Methods("PUT")
	router.HandleFunc("/student/{studentId}", controllers.DeleteStudent).Methods("DELETE")
}
