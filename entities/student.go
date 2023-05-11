package entities

import "time"

type Student struct {
	Id        string
	FullName  string
	SIN       int
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
