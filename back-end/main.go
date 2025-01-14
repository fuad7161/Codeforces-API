package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gitgub.com/fuad7161/utils"
)

func main() {
	http.HandleFunc("GET /user-info", userInfo)
	fmt.Println("Server is running at :8080")
	http.ListenAndServe(":8080", nil)
}

func userInfo(w http.ResponseWriter, r *http.Request) {
	handles := "Fuadul"
	response, err := utils.FetchUserInfo(handles)
	if err != nil {
		log.Fatalf("Error fetching user info: %v", err)
	}
	fmt.Println("Process completed successfully")
	json.NewEncoder(w).Encode(response)
}
