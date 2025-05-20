package repository

import (
	"github.com/google/uuid"
	"github.com/oceanboys24/auth-tagifast/internal/auth/dto"
	"github.com/oceanboys24/auth-tagifast/internal/auth/model"
	"gorm.io/gorm"
)

type IAuthRepository interface {
	GetUserByEmail(email string) (*model.User, error)
	GetUserById(id string) (*model.User, error)
	CreateUser(user dto.UserRegisterDTO) (*model.User, error)
}

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User

	if err := r.db.Where(&model.User{Email: email}).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *AuthRepository) GetUserById(id uuid.UUID) (*model.User, error) {
	var user model.User

	if err := r.db.Where(&model.User{ID: id}).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *AuthRepository) CreateUser(user model.User) (*model.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
