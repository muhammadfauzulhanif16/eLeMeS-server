package services

import (
	"eLeMeS/entities"
	"eLeMeS/inputs"
	"eLeMeS/repositories"
	"github.com/google/uuid"
	"time"
)

type StudentService struct {
	r repositories.Student
}

func NewStudent(r repositories.Student) *StudentService {
	return &StudentService{r}
}

type Student interface {
	Add(input inputs.AddStudent) (entities.Student, error)
}

func (s *StudentService) Add(input inputs.AddStudent) (entities.Student, error) {
	data, err := s.r.Add(
		entities.Student{
			Id:        uuid.NewString(),
			FullName:  input.FullName,
			SIN:       input.SIN,
			Password:  input.Password,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})

	if err != nil {
		return data, err
	}

	return data, nil
}
