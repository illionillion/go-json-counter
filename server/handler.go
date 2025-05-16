package server

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/illionillion/go-json-counter/utils"
)

func setCommonHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
}

func CounterHandler(w http.ResponseWriter, r *http.Request) {
	// 対象パスチェック
	path := r.URL.Path
	// もしnameのパスパラメタがある場合は取得
	name := r.PathValue("name")
	if path != "/" && !strings.HasPrefix(path, "/user/") || strings.HasPrefix(path, "/user/") && name == "" {
		slog.Warn("Invalid path or missing name parameter", "path", path, "name", name)
		http.NotFound(w, r)
		return
	}

	// 実行ファイルのパスを取得
	filePath, err := utils.GetDataFilePath()
	if err != nil {
		slog.Error("Failed to get data.json path", "error", err)
		http.Error(w, "Failed to get data.json path", http.StatusInternalServerError)
		return
	}
	slog.Info("data.json path resolved", "path", filePath)

	// data.jsonを読み込む
	counter, err := utils.ReadCounter(filePath)
	if err != nil {
		slog.Error("Failed to read data.json", "error", err)
		http.Error(w, "Failed to read data.json", http.StatusInternalServerError)
		return
	}
	slog.Info("Successfully read data.json", "counter", counter)

	if name != "" {
		counter.IncrementByName(name)
		slog.Info("Incremented counter by name", "name", name, "newCount", counter.Count)
	} else {
		counter.IncrementTotal()
		slog.Info("Incremented total counter", "newCount", counter.Count)
	}

	// data.jsonに書き込む
	jsonData, err := utils.WriteCounter(filePath, counter)
	if err != nil {
		slog.Error("Failed to write data.json", "error", err)
		http.Error(w, "Failed to write data.json", http.StatusInternalServerError)
		return
	}
	slog.Info("Successfully wrote data.json", "filePath", filePath)

	// レスポンスを返す
	setCommonHeaders(w)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
	slog.Info("Response sent", "status", http.StatusOK, "response", string(jsonData))
}
