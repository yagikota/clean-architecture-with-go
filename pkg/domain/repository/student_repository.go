package repository

import (
	"context"

	"github.com/yagikota/clean_architecture_wtih_go/pkg/domain/model"
)

// IHogeHoge represents interface of HogeHoge
type IStudentRepository interface {
	SelectAllStudents(ctx context.Context) (model.StudentSlice, error)
}
