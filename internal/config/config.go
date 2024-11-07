package config

import (
	"encoding/json"
	"os"
	"time"
)

type Config struct {
	// Host is the IP address the server will listen on
	Host string `json:"host"`
	// Port is the port the server will listen on
	Port int `json:"port"`
	// MaxUploadSize is the maximum size of an uploaded file in bytes
	MaxUploadSize int `json:"maxUploadSize"`
	// RateLimitConfig is the configuration for rate limiting
	RateLimitConfig RateLimitConfig `json:"rateLimitConfig"`
	// Storage is the type of storage to use
	StorageType    string                 `json:"storageType"`
	StorageOptions map[string]interface{} `json:"storageOptions"`
	// KeyLength is the length of the random keys
	KeyLength    int    `json:"keyLength"`
	KeyNamespace string `json:"keyNamespace"`
}

type RateLimitConfig struct {
	// Max is the maximum number of requests a client can make in the duration
	Max uint `json:"max"`
	// Duration is the time window for the rate limit
	Duration time.Duration `json:"duration"`
}

func (c *Config) Save() error {
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	return os.WriteFile("config.json", data, 0644)
}

func LoadOrSaveDefault() Config {
	data, err := os.ReadFile("config.json")
	if err != nil {
		config := DefaultConfig()
		err = config.Save()
		if err != nil {
			panic(err)
		}
		return config
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	return config
}

func DefaultConfig() Config {
	return Config{
		Host:          "0.0.0.0",
		Port:          8080,
		MaxUploadSize: 2 << 20, // 2 MB
		RateLimitConfig: RateLimitConfig{
			Max:      100,
			Duration: time.Minute,
		},
		StorageType:    "file",
		StorageOptions: map[string]interface{}{"storage-path": "data"},
		KeyLength:      10,
		KeyNamespace:   "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
	}
}
