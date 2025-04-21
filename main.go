package main

import (
	"fmt"
	"grupo35-video-auth/internal/gateway"
	"grupo35-video-auth/internal/handlers"
	"log"
	"net/http"
)

func main() {

	gateway.Oauth2Config.Init()

	http.HandleFunc("/", handlers.HandleHome)
	http.HandleFunc("/login", handlers.HandleLogin)
	http.HandleFunc("/callback", handlers.HandleCallback)
	http.HandleFunc("/logout", handlers.HandleLogout)

	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
