package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// 構造体の定義変更（外側と分離）
type NameCount struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type Counter struct {
	Count int         `json:"count"`
	Data  []NameCount `json:"data"`
}

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

func setCommonHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
}

func counterHandler(w http.ResponseWriter, r *http.Request) {

	// 対象パスチェック
	path := r.URL.Path
	if path != "/" && !strings.HasPrefix(path, "/user/") {
		http.NotFound(w, r)
		return
	}

	// 実行ファイルのパスを取得
	filePath, err := getDataFilePath()
	if err != nil {
		http.Error(w, "Failed to get data.json path", http.StatusInternalServerError)
		return
	}
	fmt.Println("data.json path: ", filePath)
	// data.jsonを読み込む
	counter, err := readCounter(filePath)
	if err != nil {
		http.Error(w, "Failed to read data.json", http.StatusInternalServerError)
		return
	}
	// もしnameのパスパラメタがある場合は取得
	name := r.PathValue("name")
	if name != "" {
		fmt.Println("name: ", name)
		// 名前の処理
		found := false
		for i, data := range counter.Data {
			if data.Name == name {
				counter.Data[i].Count++
				found = true
				break
			}
		}
		if !found && name != "" {
			counter.Data = append(counter.Data, NameCount{Name: name, Count: 1})
		}

	} else {
		// countの値を1増やす
		counter.Count++
	}

	// data.jsonに書き込む
	jsonData, err := writeCounter(filePath, counter)
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

func main() {
	http.HandleFunc("/", counterHandler)
	http.HandleFunc("/user/{name}", counterHandler)
	fmt.Println("Server Start Up........")
	port := os.Getenv("PORT")
	fmt.Println("PORT: ", port)
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe("localhost:"+port, nil))
}
