package repository

import (
	"context"

	"github.com/yagikota/clean_architecture_wtih_go/pkg/domain/model"
)

// Iはinterfaceを表現
type IStudentRepository interface {
	SelectAllStudents(ctx context.Context) (model.StudentSlice, error)
	SelectStudentByID(ctx context.Context, id int) (*model.Student, error)
}
