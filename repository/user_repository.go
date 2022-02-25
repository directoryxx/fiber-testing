package repository

import (
	"gorm.io/gorm"
	"rest-api/domain"
)

type UserRepository struct{
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) Create(role *domain.User) *domain.User {
	panic("implement me")
}

func (r *UserRepository) Update(role *domain.User, userid int) *domain.User {
	panic("implement me")
}

func (r *UserRepository) FindById(roleid int) *domain.User {
	panic("implement me")
}

func (r *UserRepository) FindAll() *[]domain.User {
	panic("implement me")
}

func (r *UserRepository) Delete(roleid int) bool {
	panic("implement me")
}

