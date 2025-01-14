package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func SaveToFile(filename string, data map[string]interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}
	return nil
}
