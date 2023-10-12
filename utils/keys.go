package utils

import "time"

func Base62() string {
	id := time.Now().Unix()
	var baseChars = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	var result []byte
	for id > 0 {
		idx := id % 62
		result = append([]byte{baseChars[idx]}, result...)
		id /= 62
	}
	return string(result)
}
