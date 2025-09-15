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

	// Same hash as ToHash256, but takes a string as input
	hs := ToHash256String("password")

	// Returns same output as ToHash256
	fmt.Printf("Hash as bytes: %v\n", hs)
}
