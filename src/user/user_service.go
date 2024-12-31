package user

import (
	"context"
	"log"
)

type UserService interface {
	AddUser(UserInput, context.Context) (User, error)
	GetWelcomeMessage(string, context.Context) (string, error)
}

type userService struct {
	userRepo UserRepo
}

func NewUserService(userRepo UserRepo) UserService {
	return &userService{userRepo: userRepo}
}

func (service *userService) AddUser(userInput UserInput, ctx context.Context) (User, error) {
	user, err := service.userRepo.addUser(userInput, ctx)
	if err != nil {
		log.Fatal(err)
	}
	return user, err
}

func (service *userService) GetWelcomeMessage(id string, ctx context.Context) (string, error) {

	user, err := service.userRepo.getUser(id, ctx)

	if err != nil {
		log.Fatal(err)
	}
	message, err := buildWelcomeMessage(user)
	if err != nil {
		log.Fatal(err)
	}
	return message, err
}

func buildWelcomeMessage(user User) (string, error) {
	return "Welcome back " + user.name + "my " + user.nationality + " friend", nil
}
