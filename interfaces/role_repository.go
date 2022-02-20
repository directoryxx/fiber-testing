package interfaces

import (
	"rest-api/domain"
)

type RoleRepository interface {
	Create(role *domain.Role) *domain.Role
	Update(role *domain.Role, roleid int) *domain.Role
	FindById(roleid int) *domain.Role
	FindAll() *[]domain.Role
	Delete(roleid int) bool
}