package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

type Counter struct {
	Count int `json:"count"`
}

func counterHandler(w http.ResponseWriter, r *http.Request) {

	// 対象パスチェック
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// 実行ファイルのパスを取得
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		http.Error(w, "Failed to get source file path", http.StatusInternalServerError)
		return
	}
	basePath := filepath.Dir(filename)
	filePath := filepath.Join(basePath, "data.json")
	fmt.Println("data.json path: ", filePath)
	// data.jsonを読み込む
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Failed to read data.json", http.StatusInternalServerError)
		return
	}
	// { count: 1 }の値を取り出す
	var counter Counter
	err = json.Unmarshal(jsonData, &counter)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusInternalServerError)
		return
	}
	// countの値を1増やす
	counter.Count++
	// data.jsonにフォーマットして書き込む
	jsonData, err = json.MarshalIndent(counter, "", "  ")
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}
	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		http.Error(w, "Failed to write data.json", http.StatusInternalServerError)
		return
	}
	// レスポンスを返す
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
	fmt.Println("Count: ", counter.Count)
	fmt.Println("data.json: ", string(jsonData))

}

func main() {
	http.HandleFunc("/", counterHandler)
	fmt.Println("Server Start Up........")
	// envファイル環境変数からポート番号を取得
	port := os.Getenv("PORT")
	fmt.Println("PORT: ", port)
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe("localhost:"+port, nil))
}
