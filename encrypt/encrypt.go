package encrypt

import (
	"security/hash"
)

func Encrypt(text, key string) string {
	output := make([]byte, len(text))
	keyBytes := hash.Hash(key)
	for i := 0; i < len(text); i++ {
		output[i] = text[i] ^ keyBytes[i%len(keyBytes)]
	}
	return string(output)
}

func Compare(hashedText, text, key string) bool {
	encrypted := Encrypt(text, key)
	return hashedText == encrypted
}
