package config

import (
	"context"
	"log"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

var (
	ClientID     = "2v0henecrkp0b6i4o4t6ougm8h"
	ClientSecret = "<your-client-secret>"
	RedirectURL  = "https://diariodonordeste.verdesmares.com.br/image/contentid/policy:1.3150265:1634763592/pug1_Easy-Resize.com.jpg"
	IssuerURL    = "https://cognito-idp.us-east-1.amazonaws.com/us-east-1_n4uOV2f8m"
	Provider     *oidc.Provider
	OAuth2Config oauth2.Config
)

func InitConfig() {
	var err error
	Provider, err = oidc.NewProvider(context.Background(), IssuerURL)
	if err != nil {
		log.Fatalf("Failed to create OIDC provider: %v", err)
	}

	OAuth2Config = oauth2.Config{
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		RedirectURL:  RedirectURL,
		Endpoint:     Provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}
}