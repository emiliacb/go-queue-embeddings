package services

import (
	"fmt"

	"github.com/emiliacb/go-queue-embeddings/internal/app/ports"
)

type EmbeddingService struct {
	embedder ports.EmbeddingAdapter
}

func NewEmbeddingService(embedder ports.EmbeddingAdapter) *EmbeddingService {
	return &EmbeddingService{
		embedder: embedder,
	}
}

func (s *EmbeddingService) EmbedOne(text string, config ports.EmbeddingConfig) ([]float32, error) {
	// TODO: Set a retry mechanism
	embedding, err := s.embedder.Embed(text, config)
	if err != nil {
		fmt.Println("Error embedding text:", err)
		return nil, err
	}

	return embedding, nil
}
