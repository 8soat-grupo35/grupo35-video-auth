package handlers

import (
	"grupo35-video-auth/internal/gateway"
	"grupo35-video-auth/internal/gateway/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"golang.org/x/oauth2"
)

func TestHandleLogin_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Mock da interface IOAuthConfig
	mockOauthConfig := mock.NewMockIOAuthConfig(ctrl)

	// Configura o comportamento esperado do mock
	expectedURL := "https://example.com/auth?state=state"
	mockOauthConfig.EXPECT().AuthCodeURL("state", oauth2.AccessTypeOffline).Return(expectedURL)

	// Substitui a configuração OAuth2 pelo mock
	gateway.Oauth2Config = mockOauthConfig

	// Cria uma requisição HTTP e um ResponseRecorder
	req := httptest.NewRequest(http.MethodGet, "/login", nil)
	rec := httptest.NewRecorder()

	// Chama a função HandleLogin
	HandleLogin(rec, req)

	// Verifica o status de redirecionamento
	if rec.Code != http.StatusFound {
		t.Errorf("esperado status %d, obtido %d", http.StatusFound, rec.Code)
	}

	// Verifica o cabeçalho de localização
	location := rec.Header().Get("Location")
	if location != expectedURL {
		t.Errorf("esperado Location %q, obtido %q", expectedURL, location)
	}
}
