package interfaces

import (
	"rest-api/api/rest/request"
	"rest-api/api/rest/response"
)

type RoleService interface {
	Create(role *request.RoleRequest) *response.RoleResponse
	Update(roleReq *request.RoleRequest,roleId int) *response.RoleResponse
	GetById(roleId int) *response.RoleResponse
	GetAll() *[]response.RoleResponse
	Delete(roleId int) bool
}
