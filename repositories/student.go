package repositories

import (
	"eLeMeS/entities"
	"gorm.io/gorm"
)

type studentRepository struct {
	db *gorm.DB
}

func StudentRepository(db *gorm.DB) *studentRepository {
	return &studentRepository{db}
}

type Student interface {
	Add(data entities.Student) (entities.Student, error)
}

func (r *studentRepository) Add(data entities.Student) (entities.Student, error) {
	if err := r.db.Create(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}
