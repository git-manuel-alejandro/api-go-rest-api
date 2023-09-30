package user

import (
	"log"
)

type Service interface {
	Create(firstName, lastName, email, phone string) error
}

type service struct {
	log  *log.Logger
	repo Repository
}

func NewService(log *log.Logger, repo Repository) Service {

	return &service{
		log:  log,
		repo: repo,
	}

}

func (s service) Create(firstName, lastName, email, phone string) error {
	s.log.Println("create user service")
	s.repo.Create(&User{})

	return nil

}
