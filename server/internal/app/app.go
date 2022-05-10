package app

import (
	"fmt"
	"github.com/Hermes-Bird/job-accelerator/internal/config"
	"github.com/Hermes-Bird/job-accelerator/internal/domain"
	"github.com/Hermes-Bird/job-accelerator/internal/handlers"
	"github.com/Hermes-Bird/job-accelerator/internal/middlewares"
	"github.com/Hermes-Bird/job-accelerator/internal/repositories"
	"github.com/Hermes-Bird/job-accelerator/internal/services"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func Start() {
	cfg := config.GetConfig()

	app := fiber.New(fiber.Config{
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  20 * time.Second,
	})

	app.Use(logger.New(logger.ConfigDefault))

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/jobs?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DbUsername,
		cfg.DbPassword,
		cfg.DbAddress,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Error while connecting to db: ", err)
	}

	err = db.AutoMigrate(
		&domain.Employee{},
		&domain.EmployeeJobDescription{},
		&domain.EmployeeEducation{},
		&domain.Language{},
		&domain.Region{},
		&domain.KeySkill{},
		&domain.Company{},
		&domain.Vacancy{},
	)

	if err != nil {
		log.Fatalln("Error while migrating db : ", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		DB:       cfg.RedisDb,
		Password: cfg.RedisPassword,
	})

	tokenRepo := repositories.NewRefreshTokenRepo(redisClient)
	empRepo := repositories.NewEmployeeRepository(db)
	companyRepo := repositories.NewCompanyRepository(db)

	authService := services.NewJwtServiceImpl(cfg.AccessSecret, cfg.RefreshSecret, tokenRepo)
	employeeService := services.NewEmployeeService(empRepo, authService)
	companyService := services.NewCompanyService(companyRepo, authService)

	authCompanyMiddleware := middlewares.NewAuthMiddleware(authService, domain.CompanyType)
	//authEmployeeMiddleware := middlewares.NewAuthMiddleware(authService, domain.EmployeeType)

	apiGroup := app.Group("/api")

	ah := handlers.NewAuthHandler(
		apiGroup,
		authService,
		employeeService,
		companyService,
		cfg.AccessTimeout,
		cfg.RefreshTimeout,
	)
	eh := handlers.NewEmployeeHandler(
		apiGroup,
		employeeService,
	)
	ch := handlers.NewCompanyHandler(
		apiGroup,
		companyService,
		authCompanyMiddleware,
	)

	ah.SetupRoutes()
	eh.SetupRoutes()
	ch.SetupRoutes()

	for _, routers := range app.Stack() {
		for _, router := range routers {
			log.Println(router.Path, router.Method)
		}
	}

	err = app.Listen(cfg.Port)
	if err != nil {
		panic(err)
	}
}
