package auth

import (
	"adv/configs"
	"adv/pkg/res"
	"encoding/json"
	"fmt"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
}
type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(route *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	route.HandleFunc("POST /auth/login", handler.Login())
	route.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		secret := handler.Config.Auth.Secret
		fmt.Println("secret", secret)
		var payload LoginRequest
		err := json.NewDecoder(req.Body).Decode(&payload)

		if err != nil {
			res.Json(w, err.Error(), http.StatusBadRequest)
		}

		fmt.Println("payload", payload)
		data := LoginResponse{
			Token: "123",
		}

		res.Json(w, data, http.StatusOK)
	}
}
func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Register World!")
	}
}
