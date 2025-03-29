package handlers

import (
	"grupo35-video-auth/internal/config"
	"net/http"

	"golang.org/x/oauth2"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	state := "state" // Replace with a secure random string in production
	url := config.OAuth2Config.AuthCodeURL(state, oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusFound)
}