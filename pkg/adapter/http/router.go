package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	apiVersion      = "/v1"
	healthCheckRoot = "/health_check"
	// student系
	studentsAPIRoot = apiVersion + "/students"
	studentIDParam  = "student_id"
)

func InitRouter() *echo.Echo {
	router := echo.New()
	router.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	// ヘルスチェック
	healthCheckGroup := router.Group(healthCheckRoot)
	{
		relativePath := ""
		healthCheckGroup.GET(relativePath, healthCheck)
	}
	return router
}
