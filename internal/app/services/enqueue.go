package services

import (
	"github.com/google/uuid"
)

type EnqueueService struct{}

func NewEnqueueService() *EnqueueService {
	return &EnqueueService{}
}

func (s *EnqueueService) Enqueue(text string) (string, error) {
	id := uuid.New().String()

	// TODO: Add the queue logic and call to the worker

	return id, nil
}
