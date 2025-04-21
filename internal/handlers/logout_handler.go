package handlers

import "net/http"

func HandleLogout(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "/", http.StatusFound)
}
