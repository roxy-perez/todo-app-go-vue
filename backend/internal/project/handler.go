package project

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func Routes(db *gorm.DB) chi.Router {
	r := chi.NewRouter()

	// List all projects
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		var projects []Project
		db.Find(&projects)
		json.NewEncoder(w).Encode(projects)
	})

	// Create a new project
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		var p Project
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		db.Create(&p)
		json.NewEncoder(w).Encode(p)
	})

	// Get one project
	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		var p Project
		if err := db.First(&p, id).Error; err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(p)
	})

	// Update a project
	r.Put("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		var p Project
		if err := db.First(&p, id).Error; err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		db.Save(&p)
		json.NewEncoder(w).Encode(p)
	})

	// Delete a project
	r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		if err := db.Delete(&Project{}, id).Error; err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})

	return r
}
