package golang_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func FindMin[T interface{ int | int64 | float64 }](first, second T) T {
	if first < second {
		return first
	}

	return second
}

func TestFindMin(t *testing.T) {
	assert.Equal(t, 100, FindMin[int](100, 101))
	assert.Equal(t, int64(100), FindMin[int64](100, 101))
	assert.Equal(t, float64(100), FindMin[float64](100, 100.1))
}

// ini contoh di Generic di Type Parameter -----------------------------------------------------------------------------

func GetFirst[T []E, E any](data T) E {
	return data[0]
}

func TestGetFirst(t *testing.T) {
	names := []string{
		"Styephen", "William", "Buli",
	}

	assert.Equal(t, "Styephen", GetFirst[[]string, string](names))
}
