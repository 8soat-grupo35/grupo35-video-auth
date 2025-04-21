package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleLogout(t *testing.T) {
	tests := []struct {
		name             string
		expectedStatus   int
		expectedLocation string
	}{
		{
			name:             "redireciona para a página inicial com sucesso",
			expectedStatus:   http.StatusFound,
			expectedLocation: "/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Cria uma requisição HTTP e um ResponseRecorder
			req := httptest.NewRequest(http.MethodGet, "/logout", nil)
			rec := httptest.NewRecorder()

			// Chama a função HandleLogout
			HandleLogout(rec, req)

			// Verifica o status de redirecionamento
			if rec.Code != tt.expectedStatus {
				t.Errorf("esperado status %d, obtido %d", tt.expectedStatus, rec.Code)
			}

			// Verifica o cabeçalho de localização
			location := rec.Header().Get("Location")
			if location != tt.expectedLocation {
				t.Errorf("esperado Location %q, obtido %q", tt.expectedLocation, location)
			}
		})
	}
}
