FROM golang:1.24-alpine

# 必要なツール
RUN apk add --no-cache git

# 作業ディレクトリ
WORKDIR /app

# ホストのコードをマウントするので COPY は不要

# デフォルトCMD（go run main.go）
CMD ["go", "run", "main.go"]
