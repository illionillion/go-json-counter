FROM golang:1.24-alpine

# 必要なツール
RUN apk add --no-cache git

# 作業ディレクトリ
WORKDIR /app

COPY . .

# air をグローバルに入れる
ENV GOBIN=/usr/local/bin
RUN go mod tidy && \
    go install github.com/cosmtrek/air@v1.49.0

CMD ["air", "-c", ".air.toml"]
