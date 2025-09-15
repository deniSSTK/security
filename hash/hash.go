package hash

type Hash interface {
	Size() int
	BlockSize() int
}

type Hash256 interface {
	Hash
}
