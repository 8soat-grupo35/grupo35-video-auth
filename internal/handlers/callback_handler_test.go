package handlers

import (
	"context"
	"errors"
	"grupo35-video-auth/internal/gateway"
	"grupo35-video-auth/internal/gateway/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"golang.org/x/oauth2"
)

func TestHandleCallback(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockOauthConfig := mock.NewMockIOAuthConfig(ctrl)

	tests := []struct {
		name           string
		queryParams    string
		mockExchange   func(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error)
		expectedStatus int
		expectedBody   string
	}{
		//{
		//	name:        "sucesso na troca de token e extração de claims",
		//	queryParams: "code=valid",
		//	mockExchange: func(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
		//		return &oauth2.Token{AccessToken: "eyJraWQiOiJEcGxNNWNhQ2UzQUNnZmZ3aWZoZ2lQTk1PcGx0dWpoc3dGbTZUU0FRaCtzPSIsImFsZyI6IlJTMjU2In0.eyJzdWIiOiI5NGE4NDRhOC05MGQxLTcwZjEtOGMxNC0xOTVkYzQxNzZhYzkiLCJpc3MiOiJodHRwczpcL1wvY29nbml0by1pZHAudXMtZWFzdC0xLmFtYXpvbmFzLmNvbVwvdXMtZWFzdC0xX3pFQ3g1Z0FweiIsInZlcnNpb24iOjIsImNsaWVudF9pZCI6IjE2ajloOWttbWxyZWdmZGNxcjBmcWxkdjZ1Iiwib3JpZ2luX2p0aSI6ImViYTE2N2U1LTFjMmYtNDY2My04MDdiLTMxNmFjZjFhMmY1ZSIsImV2ZW50X2lkIjoiZmViMDBlZDUtZTI4NC00ZjU2LTkwZWYtMzQ2MmNkZjhmMDkyIiwidG9rZW5fdXNlIjoiYWNjZXNzIiwic2NvcGUiOiJwaG9uZSBvcGVuaWQgZW1haWwiLCJhdXRoX3RpbWUiOjE3NDUyNjg1MzQsImV4cCI6MTc0NTI3MjEzNCwiaWF0IjoxNzQ1MjY4NTM0LCJqdGkiOiIyZjI1YzJlNS1lZDA4LTRjZDItYWQ1My0xMjc5NDljNTExOWIiLCJ1c2VybmFtZSI6Ijk0YTg0NGE4LTkwZDEtNzBmMS04YzE0LTE5NWRjNDE3NmFjOSJ9.ENAtblKmKvskSMURBhEBuj2ZbRWPNbUjxsW0WE2oinX83HIva6BcVofXVkyb2w_hghDcyRIXPC8r6H0KseqZQkbnlIKAGMXNc6tdwNr9ZBq2Q1sZ2ZKlzk3BvuJ2RaUCrHxNkME2iRRTI9FcJAPTjS5V64tQXf8pZAtvhmZcl5LCZ7wV7FIl8w0RuyRZf3t2Yy_AqaPZNm6kdFq6J0y3AryqRmC9fptuvEitOEv6KykyZi9yNAc8Vyw3HDFVuiKAqKRMgmySWd_5siySZU3O6wnGDM4PoVRmam1Umn_jP-u_1FYaQRKFI1z6D_nQu1H2rUi0Jjq82A-6L_YnoJBu8g"}, nil
		//	},
		//	expectedStatus: http.StatusOK,
		//	expectedBody:   "94a844a8-90d1-70f1-8c14-195dc4176ac9", // Verifica o claim "sub"
		//},
		{
			name:        "falha na troca de token",
			queryParams: "code=invalid",
			mockExchange: func(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
				return nil, errors.New("exchange error")
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "Failed to exchange token",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Configura o mock para o método Exchange
			mockOauthConfig.EXPECT().Exchange(gomock.Any(), gomock.Any()).DoAndReturn(tt.mockExchange).AnyTimes()

			// Substitui a configuração OAuth2 pelo mock
			gateway.Oauth2Config = mockOauthConfig

			// Cria uma requisição HTTP com os parâmetros de consulta
			req := httptest.NewRequest(http.MethodGet, "/callback?"+tt.queryParams, nil)
			rec := httptest.NewRecorder()

			// Chama a função HandleCallback
			HandleCallback(rec, req)

			// Verifica o status da resposta
			if rec.Code != tt.expectedStatus {
				t.Errorf("esperado status %d, obtido %d", tt.expectedStatus, rec.Code)
			}

			// Verifica o corpo da resposta
			if tt.expectedBody != "" && !strings.Contains(rec.Body.String(), tt.expectedBody) {
				t.Errorf("esperado corpo contendo %q, obtido %q", tt.expectedBody, rec.Body.String())
			}
		})
	}
}
