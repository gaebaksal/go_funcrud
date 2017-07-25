package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func (app *App) GetAllProjects(w http.ResponseWriter, r *http.Request) {

	// respondJSON
}