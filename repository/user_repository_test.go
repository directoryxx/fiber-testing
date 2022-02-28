package repository

import (
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"os"
	"path"
	"rest-api/config"
	"rest-api/infrastructure"
	"testing"
)


type SuiteUser struct {
	suite.Suite
	DB   *gorm.DB
	Redis *redis.Client
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
	s.DB = database
	s.DB.Exec("DELETE FROM users")
	s.DB.Exec("DELETE FROM roles")
}

func (s *SuiteUser) TestUserRepository_Create() {
	//role := &domain.Role{
	//	Name: "Admin",
	//}
	//roleRepo := NewRoleRepository(s.DB)
	//roleRepo.Create(role)
	//assert.NotNil(s.T(), role.ID)
	//user := &domain.User{
	//	Name: "Admin",
	//	Username: "Admin",
	//	Password: "Password",
	//	RoleID: role.ID,
	//}
	//userRepo := NewUserRepository(s.DB,s.Redis)
	//userCreate := userRepo.Create(user)
	//assert.NotNil(s.T(), userCreate.ID)
}

func (s *SuiteUser) TestUserRepository_Delete() {

}

func (s *SuiteUser) TestUserRepository_FindAll() {

}

func (s *SuiteUser) TestUserRepository_FindById() {

}

func (s *SuiteUser) TestUserRepository_Update() {

}
