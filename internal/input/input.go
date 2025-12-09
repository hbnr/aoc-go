package input

import (
	"os"
	"strings"
)

// ReadAll reads the file and returns all its content
func ReadAll(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// ReadLines reads the file and splits the content on linebreak (\n)
func ReadLines(fileName string) ([]string, error) {
	data, err := ReadAll(fileName)
	if err != nil {
		return nil, err
	}

	return strings.Split(data, "\n"), nil
}

// MustReadAll reads the file and returns its content or panics if it cannot
func MustReadAll(fileName string) string {
	data, err := ReadAll(fileName)
	if err != nil {
		panic(err)
	}

	return data
}

// MustReadLines returns a slice of each line in the file or panics
func MustReadLines(fileName string) []string {
	lines, err := ReadLines(fileName)
	if err != nil {
		panic(err)
	}

	return lines
}

// GetFileName gets the sample.txt or input.txt based on your arguments
func GetFileName(year, day string, sample bool) string {
	if len(day) == 1 {
		day = "0" + day
	}

	fileName := "/input.txt"
	if sample {
		fileName = "/sample.txt"
	}

	return "cmd/" + year + "/day" + day + fileName
}
