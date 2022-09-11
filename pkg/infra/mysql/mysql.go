package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/yagikota/clean_architecture_wtih_go/pkg/domain/model"
	"github.com/yagikota/clean_architecture_wtih_go/pkg/domain/repository"
)

type studentRepository struct {
	DB *sql.DB
}

func NewStudentRepository(db *sql.DB) repository.IStudentRepository {
	return &studentRepository{
		DB: db,
	}
}

func (sr *studentRepository) SelectAllStudents(ctx context.Context) (model.StudentSlice, error) {
	// concrete DB operation
	return model.Students().All(ctx, sr.DB)
}

func (sr *studentRepository) SelectStudentByID(ctx context.Context, studentID int) (*model.Student, error) {
	// concrete DB operation
	whereID := fmt.Sprintf("%s = ?", model.StudentColumns.ID)
	return model.Students(
		qm.Where(whereID, studentID),
	).One(ctx, sr.DB)
}
