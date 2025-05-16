package utils

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
)

func GetDataFilePath() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		slog.Error("Failed to get runtime caller info")
		return "", fmt.Errorf("cannot get runtime caller info")
	}
	basePath := filepath.Dir(filename)
	dataFilePath := filepath.Join(basePath, "data.json")
	slog.Info("Data file path resolved", "path", dataFilePath)
	return dataFilePath, nil
}

func ReadCounter(filePath string) (Counter, error) {
	var c Counter

	data, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			slog.Warn("File does not exist, initializing new counter", "filePath", filePath)
			// ファイルがない場合、初期化して保存
			c = Counter{Count: 0, Data: []NameCount{}}
			if _, writeErr := WriteCounter(filePath, c); writeErr != nil {
				slog.Error("Failed to write new counter file", "filePath", filePath, "error", writeErr)
				return c, writeErr
			}
			slog.Info("Initialized new counter file", "filePath", filePath)
			return c, nil
		}
		slog.Error("Failed to read file", "filePath", filePath, "error", err)
		return c, err
	}

	err = json.Unmarshal(data, &c)
	if err != nil {
		slog.Error("Failed to unmarshal JSON", "filePath", filePath, "error", err)
		return c, err
	}
	slog.Info("Successfully read counter file", "filePath", filePath)
	return c, nil
}

func WriteCounter(filePath string, c Counter) ([]byte, error) {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		slog.Error("Failed to marshal counter to JSON", "error", err)
		return nil, err
	}
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		slog.Error("Failed to write file", "filePath", filePath, "error", err)
		return nil, err
	}
	slog.Info("Successfully wrote counter file", "filePath", filePath)
	return data, nil
}
