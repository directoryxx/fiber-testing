package repository

import (
	"fmt"
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
}

func TestInitUser(t *testing.T) {
	suite.Run(t, new(SuiteUser))
}

func (s *SuiteUser) SetupSuite() {
	errLoadEnv := godotenv.Load(path.Join(os.Getenv("HOME")) + "/goproject/rest-api/.env")
	//helper.PanicIfError(errLoadEnv)
	config.GetConfiguration(errLoadEnv)
	dsn := config.GenerateDSNMySQL()
	database,_ := infrastructure.OpenDBMysql(dsn)
	fmt.Println(database)
	s.DB = database

}

func (s *SuiteUser) TestUserRepository_Create() {
	role := domain.Role{
		Name: "Admin",
	}
	repo := NewRoleRepository(s.DB)
	repo.Create(&role)
	assert.NotNil(s.T(), role.ID)
	user := domain.User{
		Name: "Testing",
		Username: "test",
		Password: "password",
		RoleID: role.ID,
	}
	userRepo := NewUserRepository(s.DB)
	userRepo.Create(&user)
}

func (s *SuiteUser) TestUserRepository_FindAll() {
	role := domain.Role{
		Name: "Admin",
	}
	repo := NewRoleRepository(s.DB)
	repo.Create(&role)
	assert.NotNil(s.T(), role.ID)
	user := domain.User{
		Name: "Testing",
		Username: "test",
		Password: "password",
		RoleID: role.ID,
	}
	userRepo := NewUserRepository(s.DB)
	userRepo.Create(&user)
	res := userRepo.FindAll()
	assert.NotNil(s.T(), res)
}

func (s *SuiteUser) TestUserRepository_Delete() {
	userLast := &domain.User{}
	s.DB.Model(&domain.User{}).Last(userLast)
	repo := NewUserRepository(s.DB)
	delete := repo.Delete(int(userLast.ID))
	assert.True(s.T(), delete)
}

func (s *SuiteUser) TestUserRepository_FindById() {
	userLast := &domain.User{}
	s.DB.Model(&domain.User{}).Last(userLast)
	userId := userLast.ID
	repo := NewUserRepository(s.DB)
	user := repo.FindById(int(userId))
	assert.NotNil(s.T(), user.ID)
	assert.Equal(s.T(), int(userId),int(user.ID))
}

func (s *SuiteUser) TestUserRepository_Update() {
	userLast := &domain.User{}
	s.DB.Model(&domain.User{}).Last(userLast)
	roleId := userLast.ID
	repo := NewUserRepository(s.DB)
	userCheck := domain.User{
		Name: "rubah",
	}
	repo.Update(&userCheck,int(roleId))
	role := repo.FindById(int(roleId))
	assert.Equal(s.T(), role.Name,role.Name)
}

func (s *SuiteUser) TearDownSuite() {
	s.DB.Exec("SET FOREIGN_KEY_CHECKS = 0;")
	s.DB.Exec("TRUNCATE users")
	s.DB.Exec("SET FOREIGN_KEY_CHECKS = 1;")
}