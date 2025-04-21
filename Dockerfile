FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/server ./cmd/main.go

FROM ollama/ollama:latest

RUN apt-get update && apt-get install -y --no-install-recommends supervisor curl && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/server /app/server
COPY --from=builder /app/templates/ /app/templates/
COPY --from=builder /app/static/ /app/static/

COPY supervisord.conf /etc/supervisor/conf.d/app.conf

COPY supervisord_start.sh /app/supervisord_start.sh
RUN chmod +x /app/supervisord_start.sh

EXPOSE 8080
ENTRYPOINT []
CMD ["/usr/bin/supervisord", "-c", "/etc/supervisor/conf.d/app.conf"]
