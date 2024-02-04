package utils

import (
	"encoding/json"
	"fmt"
)

func PrintJson(request interface{}) {
	jsonData, _ := json.MarshalIndent(request, "", "  ")

	// Print the JSON data to the terminal
	fmt.Println(string(jsonData))
}
