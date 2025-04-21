package domain

import (
	"sync"

	"github.com/emiliacb/go-queue-embeddings/internal/app/ports"
)

type Container struct {
	Embedder ports.EmbeddingAdapter
}

var (
	instance *Container
	once     sync.Once
)

func NewContainer(embedder ports.EmbeddingAdapter) {
	once.Do(func() {
		instance = &Container{
			Embedder: embedder,
		}
	})
}

func GetContainer() *Container {
	if instance == nil {
		panic("Container not initialized")
	}
	
	return instance
}
