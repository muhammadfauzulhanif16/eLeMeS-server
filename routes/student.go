package routes

import (
	"eLeMeS/config"
	"eLeMeS/handlers"
	"eLeMeS/repositories"
	"eLeMeS/services"
	"github.com/gin-gonic/gin"
)

var (
	db                = config.Db()
	studentRepository = repositories.StudentRepository(db)
	studentService    = services.StudentService(studentRepository)
	studentHandler    = handlers.StudentHandler(studentService)
)

func Student(r *gin.Engine) {
	v1 := r.Group("api/v1")

	student := v1.Group("/students")
	student.POST("", studentHandler.Add)
}
