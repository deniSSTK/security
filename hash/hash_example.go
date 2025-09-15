package hash

import (
	"fmt"
)

func main() {
	// Get a [32]byte hash
	h := ToHash256([]byte("password"))

	// Print the hash in different formats
	fmt.Printf("Hash as bytes: %v\n", h)             // raw byte array
	fmt.Printf("Hash as hex: %x\n", h)               // readable hex output
	fmt.Printf("Hash as string: %q\n", string(h[:])) // as string (may contain unreadable characters)
}
