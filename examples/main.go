package main

import (
	"fmt"
	"log"

	"github.com/bondzai/gogear"
)

func main() {
	// Use debugger utilities.
	data := map[string]interface{}{
		"app":   "GoGear",
		"ver":   "1.0.0",
		"notes": "composite gear example",
	}
	gogear.Print(data)

	defer gogear.TrackRuntime("main")()
	gogear.TrackRoutines()

	// Use file utilities.
	filePath := "example.txt"
	err := gogear.FileWrite(filePath, []byte("Hello, Gear!"), 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}

	content, err := gogear.FileRead(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	fmt.Printf("File Content: %s\n", content)
}
