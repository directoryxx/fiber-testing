package repository

import (
	"github.com/go-redis/redis/v8"
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


type SuiteUser struct {
	suite.Suite
	DB   *gorm.DB
	Redis *redis.Client
	RoleRepository *RoleRepository
	UserRepository *UserRepository
}

func TestUser(t *testing.T) {
	suite.Run(t, new(SuiteUser))
}

func (s *SuiteUser) SetupSuiteUser() {
	errLoadEnv := godotenv.Load(path.Join(os.Getenv("HOME")) + "/goproject/rest-api/.env")
	//helper.PanicIfError(errLoadEnv)
	config.GetConfiguration(errLoadEnv)
	dsn := config.GenerateDSNMySQL(true)
	database,_ := infrastructure.OpenDBMysql(dsn)
	s.Redis = infrastructure.OpenRedis()
	s.DB = database
	s.DB.Exec("DELETE FROM users")
	s.DB.Exec("DELETE FROM roles")
	role := &domain.Role{
		Name: "Admin",
	}
	s.RoleRepository = NewRoleRepository(s.DB)
	s.RoleRepository.Create(role)
	assert.NotNil(s.T(), role.ID)
	roleUser := &domain.Role{
		Name: "User",
	}
	s.RoleRepository.Create(roleUser)
	assert.NotNil(s.T(), roleUser.ID)
}

func (s *SuiteUser) TestUserRepository_Create() {
	user := domain.User{
		Name: "admin",
		Username: "admin",
		Password: "password",
		RoleID: uint(1),
	}
	repo := NewUserRepository(s.DB,s.Redis)
	userData := repo.Create(&user)
	assert.NotNil(s.T(), userData.ID)
}

func (s *SuiteUser) TestUserRepository_Delete() {

}

func (s *SuiteUser) TestUserRepository_FindAll() {

}

func (s *SuiteUser) TestUserRepository_FindById() {

}

func (s *SuiteUser) TestUserRepository_Update() {

}
