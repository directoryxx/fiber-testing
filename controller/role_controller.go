package controller

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"rest-api/api/rest/response"
	"rest-api/service"
)

type RoleController struct{
	Ctx *fiber.Ctx
	Service *service.RoleService
}

func NewRoleController(svc *service.RoleService) *RoleController {
	return &RoleController{
		Service: svc,
	}
}

func (r *RoleController) RoleRouter(app fiber.Router,roleSvc *service.RoleService)  {
	group := app.Group("role")
	group.Get("/", r.findAllRole())
	group.Get("/:id", r.findByIdRole())
	group.Put("/:id", r.updateRole())
	group.Delete("/:id", r.deleteRole())
	group.Post("/", r.createRole())
}

func (r *RoleController) createRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON("ok")
	}
}

func (r *RoleController) updateRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON("ok")
	}
}

func (r *RoleController) deleteRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON("ok")
	}
}

func (r *RoleController) findByIdRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON("ok")
	}
}

func (r *RoleController) findAllRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		roleAll := r.Service.GetAll()
		return c.JSON(&response.DefaultSuccess{
			Data: roleAll,
			Status: http.StatusOK,
		})
	}
}
