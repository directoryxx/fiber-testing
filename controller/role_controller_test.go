package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"net/http"
	"os"
	"path"
	"rest-api/config"
	"rest-api/helper"
	"rest-api/infrastructure"
	"rest-api/repository"
	"rest-api/service"
	"testing"
)

type Suite struct {
	suite.Suite
	RoleRepo   *repository.RoleRepository
	RoleSvc   *service.RoleService
	DB *gorm.DB
	app *fiber.App
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) SetupSuite() {
	errLoadEnv := godotenv.Load(path.Join(os.Getenv("HOME")) + "/goproject/rest-api/.env")
	helper.PanicIfError(errLoadEnv)
	dsn := config.GenerateDSNMySQL(true)
	database,_ := infrastructure.OpenDBMysql(dsn)
	s.RoleRepo = repository.NewRoleRepository(database)
	s.RoleSvc = service.NewRoleService(s.RoleRepo)
	s.DB = database
	s.app = fiber.New()
	role := NewRoleController(s.RoleSvc)
	role.RoleRouter(s.app)
}



func (s *Suite) TestRoleController_createRole() {
	values := map[string]string{"name": "test"}
	json_data, err := json.Marshal(values)
	assert.NoError(s.T(), err)
	resp, err := http.Post("http://localhost:3000/api","application/json",bytes.NewBuffer(json_data))
	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res["json"])
}

func (s *Suite) TestRoleController_deleteRole() {

}

func (s *Suite) TestRoleController_findAllRole() {

}

func (s *Suite) TestRoleController_findByIdRole() {

}

func (s *Suite) TestRoleController_updateRole() {

}
