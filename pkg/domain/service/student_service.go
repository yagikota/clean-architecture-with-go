package service

import (
	"context"

	"github.com/yagikota/clean_architecture_wtih_go/pkg/domain/model"
	"github.com/yagikota/clean_architecture_wtih_go/pkg/domain/repository"
)

type IStudentService interface {
	FindUserByUserID(ctx context.Context, studentID int) (*model.Student, error)
	FindAllUsers(ctx context.Context) (model.StudentSlice, error)
}

type studentService struct {
	studentRepository repository.IStudentRepository
}

func NewUserService(sr repository.IStudentRepository) IStudentService {
	return &studentService{
		studentRepository: sr,
	}
}

func (ss *studentService) FindUserByUserID(ctx context.Context, studentID int) (*model.Student, error) {
	return ss.studentRepository.FindUserByUserID(ctx, studentID)
}

func (ss *studentService) FindAllUsers(ctx context.Context) (model.StudentSlice, error) {
	return ss.studentRepository.FindAllUsers(ctx)
}
