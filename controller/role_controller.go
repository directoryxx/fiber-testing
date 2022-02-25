package controller

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"rest-api/api/rest/request"
	"rest-api/api/rest/response"
	"rest-api/helper"
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

func (r *RoleController) RoleRouter(app fiber.Router)  {
	group := app.Group("role")
	group.Get("/", r.findAllRole())
	group.Get("/:id", r.findByIdRole())
	group.Put("/:id", r.updateRole())
	group.Delete("/:id", r.deleteRole())
	group.Post("/", r.createRole())
}

func (r *RoleController) createRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var role *request.RoleRequest
		errRequest := c.BodyParser(&role)
		helper.PanicIfError(errRequest)

		create := r.Service.Create(role)

		return c.JSON(&response.DefaultSuccess{
			Data: create,
			Status: http.StatusOK,
		})
	}
}

func (r *RoleController) updateRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		helper.PanicIfError(err)
		var roleReq *request.RoleRequest
		errRequest := c.BodyParser(&roleReq)
		helper.PanicIfError(errRequest)
		role := r.Service.Update(roleReq,id)

		return c.JSON(&response.DefaultSuccess{
			Data: role,
			Status: http.StatusOK,
		})

	}
}

func (r *RoleController) deleteRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		helper.PanicIfError(err)
		r.Service.Delete(id)

		return c.JSON(&response.DefaultSuccess{
			Data: "true",
			Status: http.StatusOK,
		})
	}
}

func (r *RoleController) findByIdRole() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		helper.PanicIfError(err)
		role := r.Service.GetById(id)

		if (role.ID == 0){
			c.Status(http.StatusNotFound)
			return c.JSON(&response.DefaultSuccess{
				Data: nil,
				Status: http.StatusNotFound,
			})
		}

		return c.JSON(&response.DefaultSuccess{
			Data: role,
			Status: http.StatusOK,
		})
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
