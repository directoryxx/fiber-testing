package interfaces

import "rest-api/domain"

type UserRepository interface {
	Create(user *domain.User) *domain.User
	Update(user *domain.User, userid int) *domain.User
	FindById(userid int) *domain.User
	FindAll() *[]domain.User
	Delete(userid int) bool
}
