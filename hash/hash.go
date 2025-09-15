package hash

func Hash(text string) [32]byte {
	var hash [32]byte
	for i, b := range []byte(text) {
		hash[i%32] ^= b
		hash[i%32] = hash[i%32]<<5 | hash[i%32]>>4
	}
	return hash
}
