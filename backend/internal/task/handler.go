package task

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Group, db *gorm.DB) {
	// List all tasks
	e.GET("", func(c echo.Context) error {
		var tasks []Task
		if err := db.Find(&tasks).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, tasks)
	})

	// Create a task
	e.POST("", func(c echo.Context) error {
		var t Task
		if err := c.Bind(&t); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		if err := db.Create(&t).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusCreated, t)
	})

	// Get one task
	e.GET("/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		var t Task
		if err := db.First(&t, id).Error; err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
		}
		return c.JSON(http.StatusOK, t)
	})

	// Update a task
	e.PUT("/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		var t Task
		if err := db.First(&t, id).Error; err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
		}

		var input Task
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}

		t.Title = input.Title
		t.Description = input.Description
		t.Completed = input.Completed
		t.ProjectID = input.ProjectID
		db.Save(&t)

		return c.JSON(http.StatusOK, t)
	})

	// Delete a task
	e.DELETE("/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		if err := db.Delete(&Task{}, id).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.NoContent(http.StatusNoContent)
	})
}
