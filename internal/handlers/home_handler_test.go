package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleHome(t *testing.T) {
	tests := []struct {
		name           string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "renderiza a página inicial com sucesso",
			expectedStatus: http.StatusOK,
			expectedBody:   "<h1>Welcome to Cognito OIDC Go App</h1>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Cria uma requisição HTTP e um ResponseRecorder
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()

			// Chama a função HandleHome
			HandleHome(rec, req)

			// Verifica o status da resposta
			if rec.Code != tt.expectedStatus {
				t.Errorf("esperado status %d, obtido %d", tt.expectedStatus, rec.Code)
			}

			// Verifica o corpo da resposta
			if !strings.Contains(rec.Body.String(), tt.expectedBody) {
				t.Errorf("esperado corpo contendo %q, obtido %q", tt.expectedBody, rec.Body.String())
			}
		})
	}
}
