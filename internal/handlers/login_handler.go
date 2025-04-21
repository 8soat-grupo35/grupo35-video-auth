package handlers

import (
	"grupo35-video-auth/internal/gateway"
	"net/http"

	"golang.org/x/oauth2"
)

func HandleLogin(writer http.ResponseWriter, request *http.Request) {
	state := "state" // Substitua por uma string aleatória e segura em produção
	url := gateway.Oauth2Config.AuthCodeURL(state, oauth2.AccessTypeOffline)
	http.Redirect(writer, request, url, http.StatusFound)
}
