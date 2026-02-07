package storage

import (
	"encoding/json"
	"os"
	"strings"
)

type JSONSaver struct{}

type GameResult struct {
	Date     string `json:"Дата"`
	Outcome  string `json:"Исход"`
	Attempts int    `json:"Количество затраченных попыток"`
}

func SaveGameResult(filename string, result GameResult) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			data = []byte("[]")
		} else {
			return err
		}
	}

	if strings.TrimSpace(string(data)) == "" {
		data = []byte("[]")
	}

	var results []GameResult
	if err := json.Unmarshal(data, &results); err != nil {
		return err
	}

	results = append(results, result)

	out, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, out, 0644)
}
