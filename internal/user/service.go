package user

import "fmt"

type Service interface {
	Create(firstName, lastName, email, phone string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {

	return &service{
		repo: repo,
	}

}

func (s service) Create(firstName, lastName, email, phone string) error {
	fmt.Println("create user service")
	s.repo.Create(&User{})

	return nil

}
