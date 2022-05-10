package handlers

import (
	"github.com/Hermes-Bird/job-accelerator/internal/domain"
	"github.com/Hermes-Bird/job-accelerator/internal/services"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type CompanyHandler struct {
	router         fiber.Router
	companyService services.CompanyService
	authMiddleware fiber.Handler
}

func NewCompanyHandler(router fiber.Router, companyService services.CompanyService, authMiddleware fiber.Handler) *CompanyHandler {
	return &CompanyHandler{
		router:         router,
		companyService: companyService,
		authMiddleware: authMiddleware,
	}
}

func (h CompanyHandler) SetupRoutes() {
	companyGroup := h.router.Group("/company")

	companyGroup.Get("/:id", func(ctx *fiber.Ctx) error {
		id, err := strconv.Atoi(ctx.Params("id"))
		if err != nil {
			ctx.Status(http.StatusNotFound)
			return ctx.JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		company, err := h.companyService.GetCompanyById(id)

		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(company)
	})

	companyGroup.Put("/", h.authMiddleware, func(ctx *fiber.Ctx) error {
		id, ok := ctx.Locals("user-id").(int)
		if !ok {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(fiber.Map{
				"error": "wrong company id",
			})
		}

		var dto domain.UpdateCompanyDto
		err := ctx.BodyParser(&dto)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(err)
		}

		company, err := h.companyService.UpdateCompany(id, dto)
		if err != nil {
			ctx.Status(fiber.StatusBadRequest)
			return ctx.JSON(err)
		}

		ctx.Status(http.StatusOK)
		return ctx.JSON(company)
	})
}
