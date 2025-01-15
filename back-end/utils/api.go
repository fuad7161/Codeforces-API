package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"gitgub.com/fuad7161/models"
)

// FetchUserInfo to fetch user information from the codeforces
func FetchUserInfo(handle string) (models.User, error) {
	apiURL := fmt.Sprintf("https://codeforces.com/api/user.info?handles=%s", handle)

	resp, err := http.Get(apiURL)
	var usr models.User
	if err != nil {
		return usr, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return usr, fmt.Errorf("failed to read response body %v", err)
	}

	var apiResponse models.APIResponse

	err = json.Unmarshal([]byte(body), &apiResponse)
	if err != nil {
		return usr, fmt.Errorf("failed to Unmarshal data")
	}

	return apiResponse.Result[0], nil
}
