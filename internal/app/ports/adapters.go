package ports

type EmbeddingAdapter interface {
	Embed(text string, config EmbeddingConfig) ([]float32, error)
}

type EmbeddingConfig struct {
	Model string
}

type DocumentStorageAdapter interface {
	Upsert(id string, status string, progress int, data []byte, metadata map[string]interface{}) error
	Get(id string) ([]byte, error)
}
