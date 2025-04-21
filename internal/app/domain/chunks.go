package domain

import "strings"

type ChunkStrategy interface {
	Chunk(text string, config ChunkConfig) ([]string, error)
}

type ChunkConfig struct {
	ChunkSize int
	ChunkOverlap int
}

type NaiveChunkStrategy struct{}

func (s *NaiveChunkStrategy) Chunk(text string, config ChunkConfig) ([]string, error) {
	return strings.Split(text, " "), nil
}

