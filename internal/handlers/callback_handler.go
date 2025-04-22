package handlers

import (
	"context"
	"fmt"
	"grupo35-video-auth/internal/gateway"
	"html/template"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

func HandleCallback(writer http.ResponseWriter, request *http.Request) {
	ctx := context.Background()
	code := request.URL.Query().Get("code")

	// Troca o código de autorização por um token usando a interface
	rawToken, err := gateway.Oauth2Config.Exchange(ctx, code)
	if err != nil {
		http.Error(writer, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	tokenString := rawToken.AccessToken

	// Analisa o token (realize a verificação da assinatura conforme necessário em produção)
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		log.Printf("Error parsing token: %v\n", err)
		return
	}

	// Verifica se o token é válido e extrai as claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(writer, "Invalid claims", http.StatusBadRequest)
		return
	}

	// Prepare data para renderização da página
	pageData := gateway.ClaimsPage{
		AccessToken: tokenString,
		Claims:      claims,
	}

	fmt.Println("Claims:", claims)

	tmpl := template.Must(template.ParseFiles("internal/templates/claims.html"))
	tmpl.Execute(writer, pageData)
}
