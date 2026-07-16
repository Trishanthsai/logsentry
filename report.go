package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func saveReport(report Report) {
	jsonData, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling report:", err)
		return
	}

	err = os.WriteFile("report.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing report:", err)
		return
	}

	fmt.Println("Report saved to report.json")
}