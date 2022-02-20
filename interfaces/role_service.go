package interfaces

import (
	"rest-api/api/rest/request"
	"rest-api/domain"
)

type RoleService interface {
	Create(role *request.RoleRequest) *domain.Role
	Update(roleReq *request.RoleRequest,roleId int) *domain.Role
	GetById(roleId int) *domain.Role
	GetAll() *[]domain.Role
	Delete(roleId int) bool
}
