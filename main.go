package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"rest-api/config"
	"rest-api/controller"
	"rest-api/helper"
	"rest-api/infrastructure"
	"rest-api/repository"
	"rest-api/service"
)

func main() {
	errLoadEnv := godotenv.Load()
	helper.PanicIfError(errLoadEnv)

	dsn := config.GenerateDSNMySQL(false)
	database,err := infrastructure.OpenDBMysql(dsn)
	helper.PanicIfError(err)
	fmt.Println(database)


	app := fiber.New()

	root := app.Group("/api")

	// Role
	repoRole := repository.NewRoleRepository(database)
	svcRole := service.NewRoleService(repoRole)
	role := controller.NewRoleController(svcRole)
	role.RoleRouter(root)

	app.Listen(":3000")
}