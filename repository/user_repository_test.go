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
	s.DB.Exec("DELETE FROM roles")
}

func (s *SuiteUser) TestNewUserRepository(t *testing.T) {

}

func (s *SuiteUser) TestUserRepository_Create(t *testing.T) {

}

func (s *SuiteUser) TestUserRepository_Delete(t *testing.T) {

}

func (s *SuiteUser) TestUserRepository_FindAll(t *testing.T) {

}

func (s *SuiteUser) TestUserRepository_FindById(t *testing.T) {

}

func (s *SuiteUser) TestUserRepository_Update(t *testing.T) {

}
