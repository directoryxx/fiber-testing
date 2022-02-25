package interfaces

import "rest-api/domain"

type UserRepository interface {
	Create(role *domain.User) *domain.User
	Update(role *domain.User, userid int) *domain.User
	FindById(roleid int) *domain.User
	FindAll() *[]domain.User
	Delete(roleid int) bool
}
