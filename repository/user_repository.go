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

func (r *UserRepository) Create(user *domain.User) *domain.User {
	r.DB.Create(&user)
	return user
}

func (r *UserRepository) Update(user *domain.User, userid int) *domain.User {
	r.DB.Model(user).Where("id = ?", userid).Updates(user)
	return user
}

func (r *UserRepository) FindById(userid int) *domain.User {
	user := &domain.User{}
	r.DB.Model(&domain.User{}).Where("id = ?", userid).First(user)
	return user
}

func (r *UserRepository) FindAll() *[]domain.User {
	var User []domain.User
	r.DB.Model(&domain.User{}).Find(&User)
	return &User
}

func (r *UserRepository) Delete(userid int) bool {
	r.DB.Delete(&domain.User{}, userid)
	return true
}

