package models

import (
	"github.com/Serenity0204/student-management-system-backend/packages/config"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

type Student struct {
	gorm.Model
	Name       string `json:"name"`
	Age        string `json:"age"`
	Gender     string `json:"gender"`
	Instructor string `json:"instructor"`
}

func init() {
	config.ConnectToDatabase()
	db = config.GetDB()
	db.AutoMigrate(&Student{})
}

func (s *Student) CreateStudent() *Student {
	db.NewRecord(s)
	db.Create(&s)
	return s
}

func GetAllStudents() []Student {
	var Students []Student
	db.Find(&Students)
	return Students
}

func GetStudentById(Id int64) (*Student, *gorm.DB) {
	var getStudent Student
	db := db.Where("ID=?", Id).Find(&getStudent)
	return &getStudent, db
}

func DeleteStudent(Id int64) Student {
	var deletedStudent Student
	db.Where("ID=?", Id).Delete(deletedStudent)
	return deletedStudent
}
