package readtext

import (
	"bufio"
	"fmt"
	"os"
)

func OpenTextFile(fileName string) []string {
	result := []string{}
	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("Error opening file: %v", err))
	}
	defer file.Close()

	// Create a new scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Add each line to result
		result = append(result, scanner.Text())
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading file: %v", err))
	}

	return result
}
