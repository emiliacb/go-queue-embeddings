package main

import (
	"github.com/emiliacb/go-queue-embeddings/internal/app"
	"github.com/emiliacb/go-queue-embeddings/internal/app/domain"
	"github.com/emiliacb/go-queue-embeddings/internal/adapters"
)

func main() {
	domain.NewContainer(adapters.NewOllamaEmbeddingAdapter())
	app.StartServer()
}
