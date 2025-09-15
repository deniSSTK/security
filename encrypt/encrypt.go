package encrypt

import (
	"security/hash"
)

func Encrypt(text, key string) string {
	output := make([]byte, len(text))
	keyBytes := hash.ToHash256([]byte(key))
	for i := 0; i < len(text); i++ {
		output[i] = text[i] ^ keyBytes[i%len(keyBytes)]
	}
	return string(output)
}

func Decrypt(text, key string) string {
	return Encrypt(text, key)
}
