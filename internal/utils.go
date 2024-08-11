package internal

import (
	"io"
	"os"
)

func WriteToFile(content, path string) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func ReadFile(path string) (string, error) {
	// Open the file for reading
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Read the file content and skip the break lines
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
