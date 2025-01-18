package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"
)

// BeginAuthHandler initiates the OAuth process
func BeginAuthHandler(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")

	// Log which provider is being used for OAuth
	fmt.Println("Initiating OAuth for provider:", provider)

	// Set provider in context for use in the callback
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	// Start the OAuth authentication process
	gothic.BeginAuthHandler(w, r)
}

// GetAuthCallBack handles the OAuth callback after successful authentication
func GetAuthCallBack(w http.ResponseWriter, r *http.Request) {
	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	// Complete the OAuth authentication and retrieve the user
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		// Log the error if authentication fails
		fmt.Println("Error completing user auth:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the user object has a non-empty email field
	if user.Email != "" {
		// Log the authenticated user's details
		fmt.Println("Authenticated user:", user)
	} else {
		// In case no valid user is returned
		fmt.Println("No user found in callback.")
	}

	// Redirect the user to the home page or a post-login page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
