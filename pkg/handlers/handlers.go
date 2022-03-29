package handlers

import (
	"net/http"

	"github.com/anujyadav001/go-course/pkg/render"
)

// home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.html")
}

// about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.html")
}
