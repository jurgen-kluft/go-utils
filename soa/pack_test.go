package soa_test

import (
	"testing"

	"github.com/jurgen-kluft/go-utils/assert"
	"github.com/jurgen-kluft/go-utils/soa"
)

func Test_Pack(t *testing.T) {
	// Arrange
	src := make([]int32, 4096)
	states := make([]uint64, len(src))
	// Act
	states[0] = 1
	states[1] = 2
	states[111] = 3
	states[999] = 4
	packed := soa.Pack(src, states)
	// Assert
	assert.That("pack should shrink [packed] to 4", t, len(packed), 4)
}
