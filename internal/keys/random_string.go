package keys

import "math/rand"

func GenerateRandomKey(length int, namespace string) string {
	key := make([]byte, length)
	for i := range key {
		key[i] = namespace[rand.Intn(len(namespace))]
	}
	return string(key)
}
