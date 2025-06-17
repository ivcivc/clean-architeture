FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN ls -laR

RUN go build -o main .

FROM alpine:latest

# Instalar dependências necessárias
RUN apk --no-cache add ca-certificates mysql-client netcat-openbsd bash

WORKDIR /app

# Copiar binário e arquivos necessários
COPY --from=builder /app/main .
COPY --from=builder /app/sql ./sql

EXPOSE 8080 50051 8081

# Iniciar aplicação diretamente
CMD ["./main"] 