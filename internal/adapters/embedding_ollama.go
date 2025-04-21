package adapters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/emiliacb/go-queue-embeddings/internal/app/ports"
)

type ollamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

type ollamaResponse struct {
	Embedding []float32 `json:"embedding"`
}

type OllamaEmbeddingAdapter struct {}

func NewOllamaEmbeddingAdapter() *OllamaEmbeddingAdapter {
	return &OllamaEmbeddingAdapter{}
}

const (
	DefaultOllamaBaseURL = "http://localhost:11434"
	DefaultOllamaModel = "snowflake-arctic-embed2:568m-l-fp16"
)

func (a *OllamaEmbeddingAdapter) Embed(text string, config ports.EmbeddingConfig) ([]float32, error) {
	model := DefaultOllamaModel
	if config.Model != "" {
		model = config.Model
	}

	reqBody := ollamaRequest{
		Model: model,
		Prompt: text,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %w", err)
	}

	baseUrl := DefaultOllamaBaseURL
	if os.Getenv("OLLAMA_BASE_URL") != "" {
		baseUrl = os.Getenv("OLLAMA_BASE_URL")
	}

	resp, err := http.Post(
		fmt.Sprintf("%s/api/embeddings", baseUrl),
		"application/json",
		bytes.NewBuffer(jsonBody),
	)

	if err != nil {
		return nil, fmt.Errorf("error making request to Ollama: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code from Ollama: %d", resp.StatusCode)
	}

	var ollamaResp ollamaResponse
	if err := json.NewDecoder(resp.Body).Decode(&ollamaResp); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return ollamaResp.Embedding, nil
}
