package hash

const hash256StateSize = 32

const hash256BlockSize = 64

func ToHash256(data []byte) [hash256StateSize]byte {
	var state [hash256StateSize]byte

	for i := 0; i < len(data); i += hash256BlockSize {
		end := i + hash256BlockSize
		if end > len(data) {
			end = len(data)
		}
		block := data[i:end]

		for j, b := range block {
			idx := j % 32
			state[idx] ^= b
			state[idx] = state[idx]<<5 | state[idx]>>3
			if idx > 0 {
				state[idx] ^= state[idx-1]
			}
		}
	}

	for i := range state {
		state[i] ^= byte(len(data) >> (i % 8))
		state[i] = state[i]<<3 | state[i]>>5
	}

	return state
}

func Compare(hashedText [32]byte, text string) bool {
	return hashedText == ToHash256([]byte(text))
}
