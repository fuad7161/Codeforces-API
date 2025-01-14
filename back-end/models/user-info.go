package models

// User to store all user information
type User struct {
	Avatar        string `json:"avatar"`
	City          string `json:"city"`
	Country       string `json:"country"`
	Email         string `json:"email"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Handle        string `json:"handle"`
	MaxRank       string `json:"maxRank"`
	MaxRating     int    `json:"maxRating"`
	Organization  string `json:"organization"`
	Rank          string `json:"rank"`
	Rating        int    `json:"rating"`
	FriendOfCount int    `json:"friendOfCount"`
}

// UserInfo to store all the user data
var UserInfo struct {
	Result User   `json:"result"`
	Status string `json:"status"`
}

// PageData represents the data sent to the template
type PageData struct {
	User User
}
