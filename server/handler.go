package server

import (
	"fmt"
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
		http.NotFound(w, r)
		return
	}

	// 実行ファイルのパスを取得
	filePath, err := utils.GetDataFilePath()
	if err != nil {
		http.Error(w, "Failed to get data.json path", http.StatusInternalServerError)
		return
	}
	fmt.Println("data.json path: ", filePath)
	// data.jsonを読み込む
	counter, err := utils.ReadCounter(filePath)
	if err != nil {
		http.Error(w, "Failed to read data.json", http.StatusInternalServerError)
		return
	}

	if name != "" {
		counter.IncrementByName(name)
	} else {
		counter.IncrementTotal()
	}

	// data.jsonに書き込む
	jsonData, err := utils.WriteCounter(filePath, counter)
	if err != nil {
		http.Error(w, "Failed to write data.json", http.StatusInternalServerError)
		return
	}
	// レスポンスを返す
	setCommonHeaders(w)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
	fmt.Println("Count: ", counter.Count)
	fmt.Println("data.json: ", string(jsonData))
}
