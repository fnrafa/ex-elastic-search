package utils

import (
	"encoding/json"
	"log"
	"os"
)

func ReadJSON[T any](path string) T {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error reading JSON file at %s: %s", path, err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var data T
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		log.Fatalf("Error decoding JSON file at %s: %s", path, err)
	}
	return data
}
