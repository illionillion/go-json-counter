# 📊 go-json-counter

`go-json-counter` は、**Go製のシンプルなカウントAPIサーバー**です。  
アクセスに応じて JSON ファイル (`data.json`) にカウントを記録・更新します。

---

## 🔧 機能

- `/` にアクセスすると、全体のカウントが `+1`
- `/user/{name}` にアクセスすると、指定した `name` に対応するカウントが `+1`
- カウント結果は `utils/data.json` に保存され、サーバーを再起動しても保持されます

---

## ▶️ 実行方法（ローカル）

```sh
go mod tidy
go run ./main.go

# 任意のポートを指定（デフォルトは 8080）
PORT=3000 go run ./main.go
````

---

## 🐳 Docker を使った実行

### 1. `.env` ファイルの作成

```env
PORT=8080
```

### 2. コンテナをビルド・起動

```sh
docker compose up -d --build
```

### 3. API アクセス例

```sh
curl http://localhost:8080/           # 全体カウント +1
curl http://localhost:8080/user/alice # alice のカウント +1
```

---

## 📁 ディレクトリ構成

```
.
├── main.go                  # エントリーポイント
├── utils/
│   ├── counter.go           # カウントロジック
│   ├── file-util.go         # JSON 読み書き
│   └── data.json            # カウント保存ファイル
├── server/
│   └── handler.go           # HTTP ハンドラ
├── Dockerfile               # Docker ビルド設定
├── compose.yml              # Docker Compose 設定
├── .env                     # ポート指定など
└── README.md
```

---

## ⚙️ 開発用ホットリロード（Air）

ホットリロード用に [`cosmtrek/air`](https://github.com/cosmtrek/air) を使用しています。

### 1. Docker でホットリロード起動

```sh
docker compose -f compose.dev.yml up
```

※ `compose.dev.yml` には `air` を使った開発用構成が記述されています。

---

## 🧪 テスト実行

```sh
go test ./...
```

---

## 📝 ライセンス

MIT