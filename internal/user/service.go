package user

import "fmt"

type Service interface {
	Create(firstName, lastName, email, phone string) error
}

type service struct {
}

func NewService() Service {

	return &service{}

}

func (s service) Create(firstName, lastName, email, phone string) error {
	fmt.Println("create user service")

	return nil

}
