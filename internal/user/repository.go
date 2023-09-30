package user

import (
	"log"

	"gorm.io/gorm"
)

type Repository interface {
	Create(user *User) error
}

type repo struct {
	log *log.Logger
	db  *gorm.DB
}

func NewRepo(log *log.Logger, db *gorm.DB) Repository {
	return &repo{
		log: log,
		db:  db,
	}
}

func (repo *repo) Create(user *User) error {

	repo.log.Println("create user repository")
	return nil

}
