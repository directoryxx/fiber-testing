package repository

import (
	"gorm.io/gorm"
	"rest-api/domain"
)

type RoleRepository struct{
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{
		DB: db,
	}
}

func (r *RoleRepository) Create(role *domain.Role) *domain.Role {
	r.DB.Create(&role)
	return role
}

func (r *RoleRepository) Update(role *domain.Role, roleid int) *domain.Role {
	r.DB.Model(role).Where("id = ?", roleid).Updates(role)
	return role
}

func (r *RoleRepository) FindById(roleid int) *domain.Role {
	role := &domain.Role{}
	r.DB.Model(&domain.Role{}).Where("id = ?", roleid).First(role)
	return role
}

func (r *RoleRepository) FindAll() *[]domain.Role {
	var Role *[]domain.Role
	r.DB.Model(&domain.Role{}).Find(&Role)
	return Role
}

func (r *RoleRepository) Delete(roleid int) bool {
	r.DB.Delete(&domain.Role{}, roleid)
	return true
}
