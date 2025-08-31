package project

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Group, db *gorm.DB) {
	// List all projects
	e.GET("", func(c echo.Context) error {
		var projects []Project
		if err := db.Find(&projects).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, projects)
	})

	// Create a new project
	e.POST("", func(c echo.Context) error {
		var p Project
		if err := c.Bind(&p); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		if err := db.Create(&p).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusCreated, p)
	})

	// Get one project
	e.GET("/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		var p Project
		if err := db.First(&p, id).Error; err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "project not found"})
		}
		return c.JSON(http.StatusOK, p)
	})

	// Update a project
	e.PUT("/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		var p Project
		if err := db.First(&p, id).Error; err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "project not found"})
		}

		var input Project
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}

		p.Name = input.Name
		p.Description = input.Description
		db.Save(&p)

		return c.JSON(http.StatusOK, p)
	})

	// Delete a project
	e.DELETE("/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := db.Delete(&Project{}, id).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.NoContent(http.StatusNoContent)
	})
}
