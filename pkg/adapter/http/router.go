package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yagikota/clean_architecture_wtih_go/pkg/domain/service"
	"github.com/yagikota/clean_architecture_wtih_go/pkg/infra"
	"github.com/yagikota/clean_architecture_wtih_go/pkg/infra/mysql"
	"github.com/yagikota/clean_architecture_wtih_go/pkg/usecase"
)

const (
	apiVersion      = "/v1"
	healthCheckRoot = "/health_check"
	// student系
	studentsAPIRoot = apiVersion + "/students"
	studentIDParam  = "student_id"
)

func InitRouter() *echo.Echo {
	e := echo.New()
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	// ヘルスチェック
	healthCheckGroup := e.Group(healthCheckRoot)
	{
		relativePath := ""
		healthCheckGroup.GET(relativePath, healthCheck)
	}

	// student
	mySQLConn := infra.NewMySQLConnector()
	studentRepository := mysql.NewRoomRepository(mySQLConn.Conn)
	studentService := service.NewUserService(studentRepository)
	studentUsecase := usecase.NewUserUsecase(studentService)

	studentGroup := e.Group(studentsAPIRoot)
	{
		handler := NewStudentHandler(studentUsecase)
		// v1/students
		relativePath := ""
		studentGroup.GET(relativePath, handler.FindAllStudents())
	}

	return e
}
