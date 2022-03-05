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
	//redis := infrastructure.OpenRedis()
	helper.PanicIfError(err)
	fmt.Println(database)


	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	root := app.Group("/api")

	// Role
	repoRole := repository.NewRoleRepository(database)
	svcRole := service.NewRoleService(repoRole)
	role := controller.NewRoleController(svcRole,root)
	role.RoleRouter()

	// User
	//userRepo := repository.NewUserRepository(database,redis)
	//userRepo.Create(&domain.User{
	//	Name: "Admin",
	//	Username: "admin",
	//	Password: "admin",
	//	RoleID: uint(1),
	//})

	app.Listen(":3000")
}