FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/server ./cmd/main.go

# Usar la imagen de Ollama como base
FROM ollama/ollama:latest

# Copiar el binario compilado de la aplicación Go
COPY --from=builder /app/server /app/server

# Exponer solo el puerto de la aplicación Go
EXPOSE 8080

# Crear el script de inicio
RUN echo '#!/bin/sh' > /start.sh && \
    echo 'ollama serve &' >> /start.sh && \
    echo 'echo "Waiting for Ollama to start..."' >> /start.sh && \
    echo 'sleep 5' >> /start.sh && \
    echo 'echo "Pulling snowflake-arctic-embed2 model..."' >> /start.sh && \
    echo 'ollama pull snowflake-arctic-embed2:568m-l-fp16' >> /start.sh && \
    echo 'echo "Starting Go application..."' >> /start.sh && \
    echo '/app/server' >> /start.sh && \
    chmod +x /start.sh

ENTRYPOINT ["/bin/sh"]
CMD ["/start.sh"]
