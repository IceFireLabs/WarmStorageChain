package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ResourceKeyPrefix is the prefix to retrieve all Resource
	ResourceKeyPrefix = "Resource/value/"
)

// ResourceKey returns the store key to retrieve a Resource from the index fields
func ResourceKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
