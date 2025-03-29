package handlers

import (
	"context"
	"grupo35-video-auth/internal/config"
	"html/template"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

type ClaimsPage struct {
	AccessToken string
	Claims      jwt.MapClaims
}

func HandleCallback(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	code := r.URL.Query().Get("code")
	rawToken, err := config.OAuth2Config.Exchange(ctx, code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}
	tokenString := rawToken.AccessToken
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		http.Error(w, "Error parsing token", http.StatusInternalServerError)
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		http.Error(w, "Invalid claims", http.StatusBadRequest)
		return
	}
	pageData := ClaimsPage{
		AccessToken: tokenString,
		Claims:      claims,
	}
	tmpl := `
	<html>
	<body>
		<h1>User Information</h1>
		<p><strong>Access Token:</strong> {{.AccessToken}}</p>
		<ul>
			{{range $key, $value := .Claims}}
				<li><strong>{{$key}}:</strong> {{$value}}</li>
			{{end}}
		</ul>
		<a href="/logout">Logout</a>
	</body>
	</html>`
	t := template.Must(template.New("claims").Parse(tmpl))
	t.Execute(w, pageData)
}