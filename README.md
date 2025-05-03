# 📊 go-json-counter

`go-json-counter` は、**Go製のシンプルなカウントAPIサーバー**です。
アクセスに応じて JSON ファイル (`data.json`) にカウントを記録・更新します。

## 🔧 機能

* `/` にアクセスすると、全体のカウントが `+1`。
* `/user/{name}` にアクセスすると、指定した `name` に対応するカウントが `+1`。
* カウント結果は `data.json` に保存され、サーバーを再起動しても保持されます。

# 実行方法

```sh
go mod tidy
go run ./main.go
# 環境変数指定（デフォルトは8080）
PORT=3000 go run ./main.go
```