package moviesubject

import (
	"gorm.io/gorm"
)

type MovieSubject struct {
	db *gorm.DB
}

func NewMovieSubject(db *gorm.DB) *MovieSubject {
	return &MovieSubject{
		db: db,
	}
}
