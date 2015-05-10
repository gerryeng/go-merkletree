package merkletree

import (
	"testing"
)

func TestHash(t *testing.T) {

	hash1 := DOUBLE_SHA256("ab")
	hash2 := DOUBLE_SHA256("cc")
	expectedHash := DOUBLE_SHA256(hash1 + hash2)

	hash := Hash([]string{"a", "b", "c"}, DOUBLE_SHA256)
	if hash != expectedHash {
		t.Fail()
	}
}