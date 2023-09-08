// FILE: pkg/utils/utilities.go ______________________________
package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// ===================================================

func SpacerFrame(spacer string) {
	result := strings.Repeat(spacer, 35)
	fmt.Println(result)
}

// ===================================================

// LogObjectAsJSON logs a given struct as a JSON object to stdout.
// It takes a logName as the first argument, which will precede the JSON output.
// The item parameter is a generic type, so this function can accept any struct.
// The function will marshal the struct into a JSON object with indentation,
// and then write it to stdout.
//
// Usage:
//  type CarInfo struct {
//      Make        string  `json:"Make"`
//      Model       string  `json:"Model"`
//      Year        int     `json:"Year"`
//      Price       float64 `json:"Price"`
//      IsAvailable bool    `json:"IsAvailable"`
//  }
//  carData := CarInfo{"Toyota", "Camry", 2022, 30000.0, true}
//  LogObjectAsJSON("CarInfo", carData)
func LogObjectAsJSON[ObjectType any](logName string, item ObjectType) {
	// Convert item struct to JSON
	jsonData, err := json.MarshalIndent(item, "", "  ")
	Catch("Error marshaling to JSON:", err)

	// Write logName and jsonData to stdout
	fmt.Print(logName + ": ")

	_, err = os.Stdout.Write(jsonData)
	fmt.Println()
	Catch("Error writing to stdout:", err)
}

// ===================================================
