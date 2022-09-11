package service

import (
	"context"

	"github.com/yagikota/clean_architecture_wtih_go/pkg/domain/model"
	"github.com/yagikota/clean_architecture_wtih_go/pkg/domain/repository"
)

// インターフェース
type IStudentService interface {
	FindAllStudents(ctx context.Context) (model.StudentSlice, error)
	FindStudentByID(ctx context.Context, id int) (*model.Student, error)
}

// インターフェースを満たすstruct
type studentService struct {
	repo repository.IStudentRepository
}

// コンストラクタ
func NewStudentService(sr repository.IStudentRepository) IStudentService {
	return &studentService{
		repo: sr,
	}
}

// インターフェースを満たすように、メソッドを作る
func (ss *studentService) FindAllStudents(ctx context.Context) (model.StudentSlice, error) {
	return ss.repo.SelectAllStudents(ctx)
}

func (ss *studentService) FindStudentByID(ctx context.Context, id int) (*model.Student, error) {
	return ss.repo.SelectStudentByID(ctx, id)
}
