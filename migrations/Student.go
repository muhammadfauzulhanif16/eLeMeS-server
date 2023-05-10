package migrations

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	FullName string
	SIDN     int
	Password string
}
