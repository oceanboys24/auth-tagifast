package service

import (
	"errors"

	"github.com/oceanboys24/auth-tagifast/internal/auth/dto"
	"github.com/oceanboys24/auth-tagifast/internal/auth/model"
	"github.com/oceanboys24/auth-tagifast/internal/auth/repository"
	jwt "github.com/oceanboys24/auth-tagifast/internal/auth/utils/jwt"
	password "github.com/oceanboys24/auth-tagifast/internal/auth/utils/password"
)

type IAuthRepository interface {
	LoginUser(user dto.UserLoginDTO) (string, error)
	RegisterUser(user dto.UserRegisterDTO) (model.User, error)
}

type AuthService struct {
	repo *repository.AuthRepository
}

func NewAuthService(repo *repository.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) RegisterUser(user dto.UserRegisterDTO) (*model.User, error) {
	// check exists email
	checkUser, _ := s.repo.GetUserByEmail(user.Email)

	if checkUser != nil {
		return nil, errors.New("email already registered")
	}

	hashedPassword, err := password.Generate(user.Password)

	if err != nil {
		return nil, err
	}

	userCreated := model.User{
		Username: user.Username,
		Email:    user.Email,
		Password: hashedPassword,
	}

	return s.repo.CreateUser(userCreated)
}

func (s *AuthService) LoginUser(user dto.UserLoginDTO) (string, error) {
	userLogin, err := s.repo.GetUserByEmail(user.Email)
	if err != nil {
		return "", errors.New("user not found")
	}

	if err = password.Verify(userLogin.Password, user.Password); err != nil {
		return "", errors.New("email or password wrong")
	}

	token := jwt.Generate(&jwt.TokenPayload{
		ID:       (userLogin.ID).String(),
		Username: userLogin.Username,
	})

	return token, nil
}
