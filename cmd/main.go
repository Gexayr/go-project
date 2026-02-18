package main

import (
	"adv/configs"
	"adv/internal/auth"
	"adv/internal/link"
	"adv/pkg/db"
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

	//Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkReposoitory,
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("localhost:8080")
	server.ListenAndServe()
}
