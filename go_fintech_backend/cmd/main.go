// FILE: cmd/main.go ______________________________
package main

import (
	"fintech/pkg/utils"
)

// ===================================================

type CarInfo struct {
	Make        string  `json:"Make"`
	Model       string  `json:"Model"`
	Year        int     `json:"Year"`
	Price       float64 `json:"Price"`
	IsAvailable bool    `json:"IsAvailable"`
}
// ===================================================

func main() {
	utils.SpacerFrame("=")
	// =========================================

	// fmt.Println("Testing fintech backend...")
	// Initialize a CarInfo struct
	carData := CarInfo{
		Make:        "Toyota",
		Model:       "Camry",
		Year:        2022,
		Price:       30000.0,
		IsAvailable: true,
	}

	utils.LogObjectAsJSON("CarInfo", carData)


	// =========================================
	utils.SpacerFrame("=")
}
// ===================================================







