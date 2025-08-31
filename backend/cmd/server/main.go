package main

import (
	"net/http"
	"os"

	"todo-app/internal/project"
	"todo-app/internal/task"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Conexión DB
	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		e.Logger.Fatal("failed to connect database: ", err)
	}

	// Migración de los modelos
	if err := project.AutoMigrate(db); err != nil {
		e.Logger.Fatal("failed to migrate database: ", err)
	}
	if err := task.AutoMigrate(db); err != nil {
		e.Logger.Fatal("failed to migrate database: ", err)
	}

	// Rutas
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	projectGroup := e.Group("/projects")
	project.RegisterRoutes(projectGroup, db)

	taskGroup := e.Group("/tasks")
	task.RegisterRoutes(taskGroup, db)

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "5000"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
