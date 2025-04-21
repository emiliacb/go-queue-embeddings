FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/server ./cmd/main.go

FROM ollama/ollama:latest

RUN apt-get update && apt-get install -y --no-install-recommends supervisor curl && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/server /app/server

COPY supervisord.conf /etc/supervisor/conf.d/app.conf

COPY supervisor_start.sh /app/supervisor_start.sh
RUN chmod +x /app/supervisor_start.sh

EXPOSE 8080
ENTRYPOINT []
CMD ["/usr/bin/supervisord", "-c", "/etc/supervisor/conf.d/app.conf"]
