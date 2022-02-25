package service

import (
	"rest-api/api/rest/request"
	"rest-api/api/rest/response"
	"rest-api/domain"
	"rest-api/repository"
)

type RoleService struct{
	RoleRepository repository.RoleRepository
}

func NewRoleService(roleRepo *repository.RoleRepository) *RoleService {
	return &RoleService{
		RoleRepository: *roleRepo,
	}
}

func (svc *RoleService) Create(role *request.RoleRequest) *response.RoleResponse {
	roleCreate := &domain.Role{
		Name: role.Name,
	}

	roleCreated := svc.RoleRepository.Create(roleCreate)
	response := &response.RoleResponse{
		ID: int(roleCreated.ID),
		Name: roleCreated.Name,
	}
	return response
}

func (svc *RoleService) Update(roleReq *request.RoleRequest,roleId int) *response.RoleResponse {
	roleUpdate := &domain.Role{
		Name: roleReq.Name,
	}

	roleUpdated := svc.RoleRepository.Update(roleUpdate, roleId)
	response := &response.RoleResponse{
		ID: roleId,
		Name: roleUpdated.Name,
	}
	return response
}

func (svc *RoleService) GetById(roleId int) *response.RoleResponse {
	role := svc.RoleRepository.FindById(roleId)
	response := &response.RoleResponse{
		ID: int(role.ID),
		Name: role.Name,
	}
	return response
}

func (svc *RoleService) GetAll() *[]response.RoleResponse {
	role := svc.RoleRepository.FindAll()
	var roleResponses []response.RoleResponse
	for _, role := range role {
		roleResponses = append(roleResponses, response.ToRoleResponse(role))
	}
	return &roleResponses
}

func (svc *RoleService) Delete(roleId int) bool {
	role := svc.RoleRepository.Delete(roleId)
	return role
}

