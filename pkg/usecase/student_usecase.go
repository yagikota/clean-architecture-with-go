package usecase

import (
	"context"

	"github.com/yagikota/clean_architecture_wtih_go/pkg/usecase/model"
	"github.com/yagikota/clean_architecture_wtih_go/pkg/domain/service"
)

type IStudentUsecase interface {
	FindAllStudents(ctx context.Context) (model.StudentSlice, error)
	FindStudentByID(ctx context.Context, id int) (*model.Student, error)
}

type studentUsecase struct {
	svc service.IStudentService
}

func NewUserUsecase(ss service.IStudentService) IStudentUsecase {
	return &studentUsecase{
		svc: ss,
	}
}

func (su *studentUsecase) FindAllStudents(ctx context.Context) (model.StudentSlice, error) {
	msSlice, err := su.svc.FindAllStudents(ctx)
	if err != nil {
		return nil, err
	}

	sSlice := make(model.StudentSlice, 0, len(msSlice))
	for _, ms := range msSlice {
		sSlice = append(sSlice, model.StudentFromDomainModel(ms))
	}

	return sSlice, nil
}

func (su *studentUsecase) FindStudentByID(ctx context.Context, id int) (*model.Student, error) {
	ms, err := su.svc.FindStudentByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return model.StudentFromDomainModel(ms), nil
}
