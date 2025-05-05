# Stage 1: Build
FROM golang:1.24.2-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# バイナリを /app/bin/app に出力（本体をマウントしても壊れないようにする）
RUN mkdir -p bin && CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o bin/app

# Stage 2: Run
FROM scratch

WORKDIR /app

# 必要なものだけコピー
COPY --from=builder /app/bin/app ./app
COPY --from=builder /app/utils/data.json ./utils/data.json

ENTRYPOINT ["./app"]
