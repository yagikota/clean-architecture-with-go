package service

import (
	"context"

	"github.com/yagikota/clean_architecture_wtih_go/pkg/domain/model"
	"github.com/yagikota/clean_architecture_wtih_go/pkg/domain/repository"
)

// interface
type IStudentService interface {
	FindAllStudents(ctx context.Context) (model.StudentSlice, error)
}

// struct that meets interface
type studentService struct {
	repo repository.IStudentRepository
}

// constructor
func NewStudentService(sr repository.IStudentRepository) IStudentService {
	return &studentService{
		repo: sr,
	}
}

// implement methods according to interface
func (ss *studentService) FindAllStudents(ctx context.Context) (model.StudentSlice, error) {
	return ss.repo.SelectAllStudents(ctx)
}
