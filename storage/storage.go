package storage

import (
	"strconv"

	"go.uber.org/zap"
)

type Storage struct {
	data map[string]string
}

func NewStorage() *Storage {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("NewStorage")
	return &Storage{data: make(map[string]string)}
}

func (s *Storage) Set(key, value string) bool {
	s.data[key] = value
	return true
}

func (s *Storage) Get(key string) *string {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	val, err := s.data[key]
	if !err {
		logger.Info("Error in get")
		return nil
	}
	return &val
}

func (s *Storage) GetKind(key string) string {
	if _, err := strconv.Atoi(s.data[key]); err != nil {
		return "S"
	} else {
		return "D"
	}
}
