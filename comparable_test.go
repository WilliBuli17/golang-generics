package golang_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func IsSame[T comparable](value1, value2 T) bool {
	if value1 == value2 {
		return true
	}

	return false
}

func TestIsSame(t *testing.T) {
	assert.True(t, IsSame[string]("Willi", "Willi"))
	assert.True(t, IsSame[int](100, 100))
	assert.True(t, IsSame[bool](false, false))
}
