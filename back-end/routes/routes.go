package routes

import (
	"net/http"

	"gitgub.com/fuad7161/handlers"
)

func Routes() {
	http.HandleFunc("GET /", handlers.Home)
	http.HandleFunc("GET /user-info", handlers.UserInfo)
	http.HandleFunc("POST /user-info", handlers.UserInfoPost)
}
