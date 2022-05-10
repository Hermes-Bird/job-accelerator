package handlers

import (
	"github.com/Hermes-Bird/job-accelerator/internal/domain"
	"github.com/Hermes-Bird/job-accelerator/internal/services"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type EmployeeHandler struct {
	router     fiber.Router
	empService services.EmployeeService
}

func NewEmployeeHandler(router fiber.Router, empService services.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		router:     router,
		empService: empService,
	}
}

func (h EmployeeHandler) SetupRoutes() {
	emplGroup := h.router.Group("/employee")
	emplGroup.Get("/:id", func(ctx *fiber.Ctx) error {
		id, err := strconv.Atoi(ctx.Params("id"))
		if err != nil {
			ctx.Status(http.StatusNotFound)
			return ctx.JSON(err)
		}
		employee, err := h.empService.GetEmployeeById(id)
		if err != nil {
			ctx.Status(http.StatusNotFound)
			return ctx.JSON(err)
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(employee)
	})
	emplGroup.Put("/:id", func(ctx *fiber.Ctx) error {
		//id, err := strconv.Atoi(ctx.Get("user-id"))
		//if err != nil {
		//	ctx.Status(http.StatusInternalServerError)
		//	return err
		//}
		id, err := strconv.Atoi(ctx.Params("id"))
		if err != nil {
			ctx.Status(http.StatusNotFound)
			return ctx.JSON(err)
		}

		dto := domain.EmployeeUpdateDto{}
		err = ctx.BodyParser(&dto)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(err)
		}

		employee, err := h.empService.UpdateEmployee(id, dto)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(err)
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(employee)
	})
}
