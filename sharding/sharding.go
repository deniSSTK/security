package sharding

import (
	"crypto/sha1"
	"encoding/binary"

	"github.com/google/uuid"
)

func ShardedUUID(id uuid.UUID, numShards int) uint64 {
	h := sha1.Sum([]byte(id.String()))

	hashInt := binary.BigEndian.Uint64(h[:8])

	return hashInt % uint64(numShards)
}
