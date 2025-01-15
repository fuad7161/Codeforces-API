package models

// User to store all user information
type User struct {
	Avatar           string `json:"avatar"`
	City             string `json:"city"`
	Contribution     int    `json:"contribution"`
	Country          string `json:"country"`
	Email            string `json:"email"`
	FirstName        string `json:"firstName"`
	FriendOfCount    int    `json:"friendOfCount"`
	Handle           string `json:"handle"`
	LastName         string `json:"lastName"`
	LastOnlineTime   int64  `json:"lastOnlineTimeSeconds"`
	MaxRank          string `json:"maxRank"`
	MaxRating        int    `json:"maxRating"`
	Organization     string `json:"organization"`
	Rank             string `json:"rank"`
	Rating           int    `json:"rating"`
	RegistrationTime int64  `json:"registrationTimeSeconds"`
	TitlePhoto       string `json:"titlePhoto"`
}

// UserInfo to store all the user data
var UserInfo struct {
	Result User   `json:"result"`
	Status string `json:"status"`
}
