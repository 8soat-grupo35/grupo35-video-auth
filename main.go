package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/coreos/go-oidc"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/oauth2"
)

type ClaimsPage struct {
	AccessToken string
	Claims      jwt.MapClaims
}

var (
	clientID     string
	clientSecret string
	redirectURL  string
	issuerURL    string
	provider     *oidc.Provider
	oauth2Config oauth2.Config
)

func init() {
	var err error

	// Obtendo as variáveis de ambiente
	clientID = os.Getenv("CLIENT_ID")
	clientSecret = os.Getenv("CLIENT_SECRET")
	redirectURL = os.Getenv("REDIRECT_URL")
	issuerURL = os.Getenv("ISSUER_URL")

	// Inicializando o provedor OIDC
	provider, err = oidc.NewProvider(context.Background(), issuerURL)
	if err != nil {
		log.Fatalf("Failed to create OIDC provider: %v", err)
	}

	// Configurando o OAuth2
	oauth2Config = oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "email", "openid", "phone"},
	}
}

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/logout", handleLogout)
	http.HandleFunc("/callback", handleCallback)

	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	html := `
        <html>
        <body>
            <h1>Welcome to Cognito OIDC Go App</h1>
            <a href="/login">Login with Cognito</a>
        </body>
        </html>`
	fmt.Fprint(w, html)
}

func handleLogin(writer http.ResponseWriter, request *http.Request) {
	state := "state" // Substitua por uma string aleatória e segura em produção
	url := oauth2Config.AuthCodeURL(state, oauth2.AccessTypeOffline)
	http.Redirect(writer, request, url, http.StatusFound)
}

func handleCallback(writer http.ResponseWriter, request *http.Request) {
	ctx := context.Background()
	code := request.URL.Query().Get("code")

	// Troca o código de autorização por um token
	rawToken, err := oauth2Config.Exchange(ctx, code)
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
	pageData := ClaimsPage{
		AccessToken: tokenString,
		Claims:      claims,
	}

	fmt.Println("Claims:", claims)

	// Define o template HTML
	tmpl := `
    <html>
        <body>
            <h1>User Information</h1>
            <h1>JWT Claims</h1>
            <p><strong>Access Token:</strong> {{.AccessToken}}</p>
            <ul>
                {{range $key, $value := .Claims}}
                    <li><strong>{{$key}}:</strong> {{$value}}</li>
                {{end}}
            </ul>
            <a href="/logout">Logout</a>
        </body>
    </html>`

	// Analisa e executa o template
	t := template.Must(template.New("claims").Parse(tmpl))
	t.Execute(writer, pageData)
}

func handleLogout(writer http.ResponseWriter, request *http.Request) {
	// Aqui, você deve limpar a sessão ou o cookie, se armazenado
	http.Redirect(writer, request, "/", http.StatusFound)
}
