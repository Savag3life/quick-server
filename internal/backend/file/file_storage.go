package file

import (
	"QuickBin/internal/backend"
	"os"
)

type FileStorage struct {
	// Path to the directory where files are stored
	storagePath string
}

func NewFileStorage(options map[string]interface{}) backend.StorageImplement {
	path := options["storage-path"].(string)
	if path == "" {
		panic("storage-path is required for file storage")
	}
	return &FileStorage{storagePath: path}
}

func (fs *FileStorage) Save(key string, data []byte) error {
	// Save the data to the file system
	return os.WriteFile(fs.storagePath+"/"+key, data, 0644)
}

func (fs *FileStorage) Get(key string) ([]byte, error) {
	// Get the data from the file system
	data, err := os.ReadFile(fs.storagePath + "/" + key)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (fs *FileStorage) Delete(key string) error {
	// Delete the data from the file system
	return os.Remove(fs.storagePath + "/" + key)
}
