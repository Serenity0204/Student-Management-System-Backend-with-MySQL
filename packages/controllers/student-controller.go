package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Serenity0204/student-management-system-backend/packages/models"
	"github.com/Serenity0204/student-management-system-backend/packages/utils"
	"github.com/gorilla/mux"
)

var NewStudent models.Student

func GetStudent(w http.ResponseWriter, req *http.Request) {
	students := models.GetAllStudents()
	res, _ := json.Marshal(students)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetStudentById(w http.ResponseWriter, req *http.Request) {
	studentHash := mux.Vars(req)
	studentId := studentHash["studentId"]
	Id, err := strconv.ParseInt(studentId, 0, 0)
	if err != nil {
		fmt.Print("error while parsing")
	}
	studentDetails, _ := models.GetStudentById(Id)
	res, _ := json.Marshal(studentDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateStudent(w http.ResponseWriter, req *http.Request) {
	student := &models.Student{}
	utils.ParseBody(req, student)
	s := student.CreateStudent()
	res, _ := json.Marshal(s)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteStudent(w http.ResponseWriter, req *http.Request) {
	studentHash := mux.Vars(req)
	studentId := studentHash["studentId"]
	Id, err := strconv.ParseInt(studentId, 0, 0)
	if err != nil {
		fmt.Print("error while parsing")
	}
	deletedStudent := models.DeleteStudent(Id)
	res, _ := json.Marshal(deletedStudent)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateStudent(w http.ResponseWriter, req *http.Request) {
	updatedStudent := &models.Student{}
	utils.ParseBody(req, updatedStudent)
	studentHash := mux.Vars(req)
	studentId := studentHash["studentId"]
	Id, err := strconv.ParseInt(studentId, 0, 0)
	if err != nil {
		fmt.Print("error while parsing")
	}
	studentDetails, db := models.GetStudentById(Id)
	if updatedStudent.Name != "" {
		studentDetails.Name = updatedStudent.Name
	}
	if updatedStudent.Age != "" {
		studentDetails.Age = updatedStudent.Age
	}
	if updatedStudent.Gender != "" {
		studentDetails.Gender = updatedStudent.Gender
	}
	if updatedStudent.Instructor != "" {
		studentDetails.Instructor = updatedStudent.Instructor
	}
	db.Save(&studentDetails)
	res, _ := json.Marshal(studentDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
