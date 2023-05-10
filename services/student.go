package services

import (
	"eLeMeS/entities"
	"eLeMeS/inputs"
	"eLeMeS/repositories"
	"github.com/google/uuid"
	"time"
)

type studentService struct {
	r repositories.Student
}

func StudentService(r repositories.Student) *studentService {
	return &studentService{r}
}

type Student interface {
	Add(input inputs.AddStudent) (entities.Student, error)
}

func (s *studentService) Add(input inputs.AddStudent) (entities.Student, error) {
	data, err := s.r.Add(
		entities.Student{
			Id:        uuid.NewString(),
			FullName:  input.FullName,
			SIDN:      input.SIDN,
			Password:  input.Password,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})

	if err != nil {
		return data, err
	}

	return data, nil
}
