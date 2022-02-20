package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"rest-api/config"
	"rest-api/helper"
	"rest-api/infrastructure"
)

func main() {

	errLoadEnv := godotenv.Load()
	helper.PanicIfError(errLoadEnv)

	dsn := config.GenerateDSNMySQL(false)
	database,err := infrastructure.OpenDBMysql(dsn)
	helper.PanicIfError(err)
	fmt.Println(database)


	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}