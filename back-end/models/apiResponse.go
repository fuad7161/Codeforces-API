package models

type APIResponse struct {
	Result []User `json:"result"`
	Status string `json:"status"`
}
