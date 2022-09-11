package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yagikota/clean_architecture_wtih_go/pkg/usecase"
)

type studentHandler struct {
	usecase usecase.IStudentUsecase
}

func NewStudentHandler(su usecase.IStudentUsecase) *studentHandler {
	return &studentHandler{
		usecase: su,
	}
}

func (sh *studentHandler) FindAllStudents() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		student, err := sh.usecase.FindAllStudents(ctx)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, student)
	}
}
