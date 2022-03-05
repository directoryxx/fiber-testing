package repository

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"os"
	"path"
	"rest-api/config"
	"rest-api/domain"
	"rest-api/infrastructure"
	"testing"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) SetupSuite() {
	errLoadEnv := godotenv.Load(path.Join(os.Getenv("HOME")) + "/goproject/rest-api/.env")
	//helper.PanicIfError(errLoadEnv)
	config.GetConfiguration(errLoadEnv)
	dsn := config.GenerateDSNMySQL()
	database,_ := infrastructure.OpenDBMysql(dsn)
	s.DB = database
	s.DB.Exec("DELETE FROM roles")
}

func (s *Suite) TestRoleRepository_Create() {
	role := &domain.Role{
		Name: "Admin",
	}
	repo := NewRoleRepository(s.DB)
	repo.Create(role)
	assert.NotNil(s.T(), role.ID)
}

func (s *Suite) TestRoleRepository_Delete() {
	roleLast := &domain.Role{}
	s.DB.Model(&domain.Role{}).Last(roleLast)
	repo := NewRoleRepository(s.DB)
	delete := repo.Delete(int(roleLast.ID))
	assert.True(s.T(), delete)
}

func (s *Suite) TestRoleRepository_FindAll() {
	repo := NewRoleRepository(s.DB)
	roleCreate := &domain.Role{
		Name: "Admin",
	}
	repo.Create(roleCreate)
	res := repo.FindAll()
	assert.NotNil(s.T(), res)
}

func (s *Suite) TestRoleRepository_FindById() {
	roleLast := &domain.Role{}
	s.DB.Model(&domain.Role{}).Last(roleLast)
	roleId := roleLast.ID
	repo := NewRoleRepository(s.DB)
	role := repo.FindById(int(roleId))
	assert.NotNil(s.T(), role.ID)
	assert.Equal(s.T(), int(roleId),int(role.ID))
}

func (s *Suite) TestRoleRepository_Update() {
	roleLast := &domain.Role{}
	s.DB.Model(&domain.Role{}).Last(roleLast)
	roleId := roleLast.ID
	repo := NewRoleRepository(s.DB)
	roleCheck := domain.Role{
		Name: "rubah",
	}
	repo.Update(&roleCheck,int(roleId))
	role := repo.FindById(int(roleId))
	assert.Equal(s.T(), role.Name,role.Name)
}

func (s *Suite) TearDownSuite() {
	s.DB.Exec("SET FOREIGN_KEY_CHECKS = 0")
	s.DB.Exec("TRUNCATE roles")
	s.DB.Exec("SET FOREIGN_KEY_CHECKS = 1")
}
