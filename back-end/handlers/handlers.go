package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gitgub.com/fuad7161/utils"
)

func UserInfo(w http.ResponseWriter, r *http.Request) {
	handles := "Fuadul"
	response, err := utils.FetchUserInfo(handles)
	if err != nil {
		log.Fatalf("Error fetching user info: %v", err)
	}
	fmt.Println("Process completed successfully")
	json.NewEncoder(w).Encode(response)
}
