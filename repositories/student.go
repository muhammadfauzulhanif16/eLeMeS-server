package repositories

import (
	"eLeMeS/entities"
	"gorm.io/gorm"
)

type StudentRepository struct {
	db *gorm.DB
}

func NewStudent(db *gorm.DB) *StudentRepository {
	return &StudentRepository{db}
}

type Student interface {
	Add(data entities.Student) (entities.Student, error)
}

func (r *StudentRepository) Add(data entities.Student) (entities.Student, error) {
	if err := r.db.Create(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}
