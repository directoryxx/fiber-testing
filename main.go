package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"os"
	"rest-api/config"
	"rest-api/controller"
	"rest-api/helper"
	"rest-api/infrastructure"
	"rest-api/repository"
	"rest-api/service"
)

func main() {
	app := SetupInit()
	app.Listen(":3000")

	if os.Getenv("TESTING") == "true" {
		app.Shutdown()
	}
}

func SetupInit() *fiber.App {
	errLoadEnv := godotenv.Load()
	config.GetConfiguration(errLoadEnv)

	dsn := config.GenerateDSNMySQL()
	database,err := infrastructure.OpenDBMysql(dsn)
	//redis := infrastructure.OpenRedis()
	helper.PanicIfError(err)
	fmt.Println(database)

	app := fiber.New()

	root := app.Group("/api")

	// Role
	repoRole := repository.NewRoleRepository(database)
	svcRole := service.NewRoleService(repoRole)
	role := controller.NewRoleController(svcRole,root)
	role.RoleRouter()

	return app
}