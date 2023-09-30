package user

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	Create(user *User) error
}

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) Repository {
	return &repo{
		db: db,
	}
}

func (repo *repo) Create(user *User) error {

	fmt.Println("create user repository")
	return nil

}
