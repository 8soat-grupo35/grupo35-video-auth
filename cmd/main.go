package main

import (
	"fmt"
	"grupo35-video-auth/internal/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HandleHome)
	http.HandleFunc("/login", handlers.HandleLogin)
	http.HandleFunc("/logout", handlers.HandleLogout)
	http.HandleFunc("/callback", handlers.HandleCallback)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}