package main

import (
	"context"

	"golang.org/x/oauth2"
)

type IOAuthConfig interface {
	AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string
	Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error)
}

type OAuth2Config struct {
	Config oauth2.Config
}

func (o *OAuth2Config) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	return o.Config.AuthCodeURL(state, opts...)
}
func (o *OAuth2Config) Exchange(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	return o.Config.Exchange(ctx, code, opts...)
}
