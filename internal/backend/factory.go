package backend

import (
	"QuickBin/internal/backend/aws"
	"QuickBin/internal/backend/file"
	"QuickBin/internal/config"
	"fmt"
)

func NewStorage(cfg config.Config) (StorageImplement, error) {
	switch cfg.StorageType {
	case "file":
		return file.NewFileStorage(cfg.StorageOptions), nil
	case "aws":
		return aws.NewAWSStorage(cfg.StorageOptions), nil
	default:
		return nil, fmt.Errorf("unknown storage type: %s", cfg.StorageType)
	}
}
