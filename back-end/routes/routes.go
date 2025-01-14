package routes

import (
	"net/http"

	"gitgub.com/fuad7161/handlers"
)

func Routes() {
	http.HandleFunc("GET /user-info", handlers.UserInfo)
}
