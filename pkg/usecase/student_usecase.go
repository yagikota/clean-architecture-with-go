package usecase

import (
	"context"

	"github.com/yagikota/clean_architecture_wtih_go/pkg/domain/model"
	"github.com/yagikota/clean_architecture_wtih_go/pkg/domain/service"
)

type IUserUsecase interface {
	FindUserByUserID(ctx context.Context, studentID int) (*model.Student, error)
	FindAllUsers(ctx context.Context) (model.StudentSlice, error)
}

type studentUsecase struct {
	studentService service.IStudentService
}

func NewUserUsecase(ss service.IStudentService) IUserUsecase {
	return &studentUsecase{
		studentService: ss,
	}
}

func (su *studentUsecase) FindUserByUserID(ctx context.Context, userID int) (*model.Student, error) {
	mu, err := su.studentService.FindUserByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return model.UserFromDomainModel(mu), err
}

func (su *studentUsecase) FindAllUsers(ctx context.Context) (model.StudentSlice, error) {
	msSlice, err := su.studentService.FindAllUsers(ctx)
	if err != nil {
		return nil, err
	}


	return uSlice, nil
}
