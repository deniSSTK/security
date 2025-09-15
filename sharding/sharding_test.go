package sharding

import (
	"testing"

	"github.com/google/uuid"
)

func TestShardedUUID(t *testing.T) {
	numShards := 4

	tests := []struct {
		name          string
		id            uuid.UUID
		expectedShard uint64
	}{
		{
			name:          "TestShardedUUID 1",
			id:            uuid.MustParse("550e8400-e29b-41d4-a716-446655440000"),
			expectedShard: 1,
		},
		{
			name:          "TestShardedUUID 2",
			id:            uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"),
			expectedShard: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shard := ShardedUUID(tt.id, numShards)

			if shard != tt.expectedShard {
				t.Errorf("expected shard %d, got %d", tt.expectedShard, shard)
			}
		})
	}
}

func TestShardedUUIDEqualValues(t *testing.T) {
	numShards := 4
	countTests := 10

	var ids []uuid.UUID
	for i := 0; i < countTests; i++ {
		ids = append(ids, uuid.New())
	}

	for _, tt := range ids {
		t.Run("Id="+tt.String(), func(t *testing.T) {
			shard1 := ShardedUUID(tt, numShards)
			shard2 := ShardedUUID(tt, numShards)

			if shard1 != shard2 {
				t.Errorf("expected shard %d, got %d", shard1, shard2)
			}
		})
	}
}
