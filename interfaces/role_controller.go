package interfaces

import (
	"github.com/gofiber/fiber/v2"
)

type RoleController interface {
	Create(c *fiber.Ctx)
	Update(c *fiber.Ctx)
	Delete(c *fiber.Ctx)
	FindById(c *fiber.Ctx)
	FindAll(c *fiber.Ctx)
}
