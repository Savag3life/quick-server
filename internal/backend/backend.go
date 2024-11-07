package backend

type StorageImplement interface {
	Save(key string, data []byte) error
	Get(key string) ([]byte, error)
	Delete(key string) error
}
