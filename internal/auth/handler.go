package auth

import (
	"adv/configs"
	"adv/pkg/req"
	"adv/pkg/res"
	"fmt"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}
type AuthHandler struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(route *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	route.HandleFunc("POST /auth/login", handler.Login())
	route.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		fmt.Println("payload", body)
		data := LoginResponse{
			Token: "123",
		}

		res.Json(w, data, http.StatusOK)
	}
}
func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}
		/*		fmt.Println("payload", body)
				data := LoginResponse{
					Token: "123",
				}

				res.Json(w, data, http.StatusOK)*/

		handler.AuthService.Register(body.Email, body.Password, body.Name)
	}
}
