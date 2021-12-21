package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		return User{}, err
	}
	user := User{
		Name:           input.Name,
		Occupation:     input.Occupation,
		Email:          input.Email,
		PasswordHashed: string(passwordHashed),
		Role:           "user",
	}
	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	user, err := s.repository.FindByEmail(input.Email)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("invalid email/password")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHashed), []byte(input.Password)); err != nil {
		return user, errors.New("invalid email/password")
	}
	return user, nil
}
