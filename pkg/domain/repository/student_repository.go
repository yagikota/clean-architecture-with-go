package repository

import (
	"context"

	"github.com/yagikota/clean_architecture_wtih_go/pkg/domain/model"
)

// Iはinterfaceを表現
type IStudentRepository interface {
	FindAllStudents(ctx context.Context) (model.StudentSlice, error)
}
