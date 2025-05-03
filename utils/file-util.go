package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func getDataFilePath() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("cannot get runtime caller info")
	}
	basePath := filepath.Dir(filename)
	return filepath.Join(basePath, "data.json"), nil
}

func readCounter(filePath string) (Counter, error) {
	var c Counter
	data, err := os.ReadFile(filePath)
	if err != nil {
		return c, err
	}
	err = json.Unmarshal(data, &c)
	return c, err
}

func writeCounter(filePath string, c Counter) ([]byte, error) {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return nil, err
	}
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return nil, err
	}
	return data, nil
}
