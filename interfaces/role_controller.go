package interfaces

import (
	"github.com/gofiber/fiber/v2"
)

type RoleController interface {
	createRole() fiber.Handler
	updateRole() fiber.Handler
	deleteRole() fiber.Handler
	findByIdRole() fiber.Handler
	findAllRole() fiber.Handler
}
