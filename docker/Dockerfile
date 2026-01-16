FROM golang:1.25-alpine AS builder
WORKDIR /app

RUN go install github.com/air-verse/air@latest

# ルートの go.mod と go.sum をコピー
COPY go.mod go.sum ./
RUN go mod download

# 全ソースコードをコピー
COPY . .

ARG APP_NAME
ENV APP_NAME=${APP_NAME}

CMD ["air", "--build.cmd", "go build -o main ./cmd/${APP_NAME}/main.go", "--build.bin", "./main"]
