package digest256

import (
	"bytes"
	"security/hash"
)

const size = 32

const blockSize = 64

type Digest256 struct {
	state [size]byte
}

func New() hash.Hash256 {
	return &Digest256{}
}

func (d *Digest256) Size() int {
	return size
}

func (d *Digest256) BlockSize() int {
	return blockSize
}

func ToHash256(data []byte) []byte {
	sum := toHash256(data)
	return sum[:]
}

func toHash256(data []byte) [size]byte {
	var state [size]byte

	for i := 0; i < len(data); i += blockSize {
		end := i + blockSize
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

func Compare(hashedText []byte, text string) bool {
	return bytes.Equal(hashedText, []byte(text))
}
