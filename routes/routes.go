package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mysterybee07/go-oauth/controllers"
)

// Route sets up the routes for the application
func Route() http.Handler {
	router := chi.NewRouter()

	// Routes for OAuth
	router.Get("/{provider}/auth", controllers.BeginAuthHandler)    // Start OAuth
	router.Get("/{provider}/callback", controllers.GetAuthCallBack) // OAuth callback

	return router
}
