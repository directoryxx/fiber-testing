package service

import (
	"rest-api/api/rest/request"
	"rest-api/domain"
	"rest-api/repository"
)

type RoleService struct{
	RoleRepository repository.RoleRepository
}

func NewRoleService(roleRepo *repository.RoleRepository) RoleService {
	return RoleService{
		RoleRepository: *roleRepo,
	}
}

func (svc *RoleService) Create(role *request.RoleRequest) *domain.Role {
	roleCreate := &domain.Role{
		Name: role.Name,
	}

	roleCreated := svc.RoleRepository.Create(roleCreate)
	return roleCreated
}

func (svc *RoleService) Update(roleReq *request.RoleRequest,roleId int) *domain.Role {
	roleUpdate := &domain.Role{
		Name: roleReq.Name,
	}

	roleUpdated := svc.RoleRepository.Update(roleUpdate, roleId)
	return roleUpdated
}

func (svc *RoleService) GetById(roleId int) *domain.Role {
	role := svc.RoleRepository.FindById(roleId)
	return role
}

func (svc *RoleService) GetAll() *[]domain.Role {
	role := svc.RoleRepository.FindAll()
	return role
}

func (svc *RoleService) Delete(roleId int) bool {
	role := svc.RoleRepository.Delete(roleId)
	return role
}

