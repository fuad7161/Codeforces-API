package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// FetchUserInfo to fetch user information from the codeforces
func FetchUserInfo(handle string) (map[string]interface{}, error) {
	apiURL := fmt.Sprintf("https://codeforces.com/api/user.info?handles=%s", handle)

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body %v", err)
	}
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to Unmarshal data")
	}
	return data, nil
}
