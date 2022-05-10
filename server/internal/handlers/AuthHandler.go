package handlers

import (
	"github.com/Hermes-Bird/job-accelerator/internal/domain"
	"github.com/Hermes-Bird/job-accelerator/internal/services"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
	"time"
)

type AuthHandler struct {
	router         fiber.Router
	companyService services.CompanyService
	authService    services.AuthService
	empService     services.EmployeeService
	accessDur      time.Duration
	refreshDur     time.Duration
}

func NewAuthHandler(
	router fiber.Router,
	as services.AuthService,
	es services.EmployeeService,
	cs services.CompanyService,
	accessDur time.Duration,
	refreshDur time.Duration) *AuthHandler {
	return &AuthHandler{
		router:         router,
		authService:    as,
		empService:     es,
		companyService: cs,
		accessDur:      accessDur,
		refreshDur:     refreshDur,
	}
}

func (h AuthHandler) SetupRoutes() {
	authGroup := h.router.Group("/auth")

	authGroup.Post("/employee/register", func(ctx *fiber.Ctx) error {
		dto := domain.CreateEmployeeDto{}
		err := ctx.BodyParser(&dto)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(err)
		}

		newEmployee, err := h.empService.CreateEmployee(dto)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(err)
		}

		tokenPair, err := h.authService.GenerateTokenPair(strconv.Itoa(newEmployee.Id), domain.EmployeeType, h.accessDur, h.refreshDur)

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(err)
		}

		ctx.Status(http.StatusAccepted)
		return ctx.JSON(tokenPair)
	})

	authGroup.Post("/employee/login", func(ctx *fiber.Ctx) error {
		dto := domain.LoginEmployeeDto{}
		err := ctx.BodyParser(&dto)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(err)
		}
		id, err := h.empService.CheckEmployeeCreds(dto)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(map[string]string{"error": err.Error()})
		}

		pair, err := h.authService.GenerateTokenPair(strconv.Itoa(id), domain.EmployeeType, h.accessDur, h.refreshDur)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(err)
		}

		ctx.Status(http.StatusAccepted)
		return ctx.JSON(pair)
	})

	authGroup.Post("company/register", func(ctx *fiber.Ctx) error {
		dto := domain.CreateCompanyDto{}
		err := ctx.BodyParser(&dto)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(err)
		}

		newEmployee, err := h.companyService.CreateCompany(dto)
		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(err)
		}

		tokenPair, err := h.authService.GenerateTokenPair(strconv.Itoa(newEmployee.Id), domain.CompanyType, h.accessDur, h.refreshDur)

		if err != nil {
			ctx.Status(http.StatusInternalServerError)
			return ctx.JSON(err)
		}

		ctx.Status(http.StatusAccepted)
		return ctx.JSON(tokenPair)
	})

	authGroup.Post("/company/login", func(ctx *fiber.Ctx) error {
		dto := domain.LoginCompanyDto{}
		err := ctx.BodyParser(&dto)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(err)
		}

		id, err := h.companyService.CheckCompanyCreds(dto)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(map[string]string{"error": err.Error()})
		}

		pair, err := h.authService.GenerateTokenPair(strconv.Itoa(id), domain.CompanyType, h.accessDur, h.refreshDur)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(err)
		}

		ctx.Status(http.StatusAccepted)
		return ctx.JSON(pair)
	})
}
