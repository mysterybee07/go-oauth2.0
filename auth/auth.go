package auth

import (
	"log"
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

const (
	key    = "randomString" // This should be a secret, ideally from an env var
	isProd = false          // Set to true in production if using HTTPS
	MaxAge = 86400 * 30     // 30 days session duration
)

func NewAuth() {
	// Get environment variables
	googleClientId := os.Getenv("GOOGLE_CLIENT_ID")
	googleClientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	redirectUrl := os.Getenv("REDIRECT_URL")

	// Check if the environment variables are set
	if googleClientId == "" || googleClientSecret == "" || redirectUrl == "" {
		log.Fatal("Missing required environment variables")
	}

	// Initialize session store
	store := sessions.NewCookieStore([]byte(key))

	// Set session options
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   MaxAge,
		HttpOnly: true,
		Secure:   isProd, // Set Secure to true in production (if using HTTPS)
	}

	// Use the session store with gothic
	gothic.Store = store

	// Set up Google OAuth provider
	goth.UseProviders(
		google.New(googleClientId, googleClientSecret, redirectUrl),
	)

	log.Println("OAuth initialized with Google provider")
}
