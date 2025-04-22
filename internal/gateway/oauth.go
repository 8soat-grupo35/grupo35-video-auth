package gateway

import (
	"context"
	"log"
	"os"

	"github.com/coreos/go-oidc"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/oauth2"
)

//go:generate mockgen -source=oauth.go -destination=mock/oauth_mock.go -package=mock
type ClaimsPage struct {
	AccessToken string
	Claims      jwt.MapClaims
}

type IOAuthConfig interface {
	AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string
	Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error)
}

type OAuth2Config struct {
	Config oauth2.Config
}

var (
	clientID     string
	clientSecret string
	redirectURL  string
	issuerURL    string
	provider     *oidc.Provider
	Oauth2Config IOAuthConfig
)

func Init() {
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

	// Configurando o OAuth2 usando OAuth2Config
	Oauth2Config = &OAuth2Config{
		Config: oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  redirectURL,
			Endpoint:     provider.Endpoint(),
			Scopes:       []string{"email", "openid", "profile"},
		},
	}
}

func (o *OAuth2Config) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	return o.Config.AuthCodeURL(state, opts...)
}
func (o *OAuth2Config) Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	return o.Config.Exchange(ctx, code, opts...)
}
