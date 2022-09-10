package repository

import (
	"context"

	"github.com/yagikota/clean_architecture_wtih_go/pkg/domain/model"
)

// Iはinterfaceを表現
type IStudentRepository interface {
	FindUserByUserID(ctx context.Context, studentID int) (*model.Student, error)
	FindAllUsers(ctx context.Context) (model.StudentSlice, error)
}
