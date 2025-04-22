package services

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Status string

const (
	StatusProcessing Status = "processing"
	StatusCompleted  Status = "completed"
	StatusFailed     Status = "failed"
)

type DocumentFile struct {
	ID     string `json:"id"`
	Status Status `json:"status"`
}

type DocumentService struct{}

func NewDocumentService() (*DocumentService, error) {
	fmt.Println("DocumentService: Using temp dir", os.TempDir())
	return &DocumentService{}, nil
}

func (s *DocumentService) getFilePath(id string) string {
	storageSubDir := os.TempDir()
	folderName := "go-queue-embeddings-documents"
	return filepath.Join(storageSubDir, folderName, id+".json")
}

func (s *DocumentService) SaveDocument(id string, status Status) error {
	filePath := s.getFilePath(id)

	data := DocumentFile{
		ID:     id,
		Status: Status(status),
	}

	// Format with human readable JSON
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal document file for id '%s': %w", id, err)
	}

	err = os.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write document file '%s': %w", filePath, err)
	}

	return nil
}

func (s *DocumentService) ReadDocument(id string) (DocumentFile, error) {
	filePath := s.getFilePath(id)
	var data DocumentFile

	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return data, fmt.Errorf("document file not found for id '%s' at '%s'", id, filePath)
		}
		return data, fmt.Errorf("failed to read document file '%s': %w", filePath, err)
	}

	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return data, fmt.Errorf("failed to unmarshal document file from file '%s': %w", filePath, err)
	}

	return data, nil
}
