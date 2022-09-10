package mysql

import (
	"context"
	"database/sql"

	"github.com/yagikota/clean_architecture_wtih_go/pkg/domain/model"
	"github.com/yagikota/clean_architecture_wtih_go/pkg/domain/repository"
)

type studentRepository struct {
	DB *sql.DB
}

func NewRoomRepository(db *sql.DB) repository.IStudentRepository {
	return &studentRepository{
		DB: db,
	}
}

func (sr *studentRepository) FindAllStudents(ctx context.Context) (model.StudentSlice, error) {
	// DBに対する具体的な操作
	return model.Students().All(ctx, sr.DB)
}
