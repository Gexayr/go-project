package main

import (
	"adv/configs"
	"adv/internal/auth"
	"adv/internal/link"
	"adv/internal/user"
	"adv/pkg/db"
	"adv/pkg/middleware"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	db := db.NewDb(conf)
	//fmt.Println(conf)
	router := http.NewServeMux()

	//Repositories
	linkReposoitory := link.NewLinkRepository(db)
	userReposoitory := user.NewUserRepository(db)

	// Services
	authService := auth.NewAuthService(userReposoitory)

	//Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})

	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkReposoitory,
	})

	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := &http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}

	fmt.Println("localhost:8080")
	server.ListenAndServe()
}
