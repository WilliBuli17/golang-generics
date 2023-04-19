package golang_generics

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestLength(t *testing.T) {
	resultString := Length[string]("Willi Buli")
	assert.Equal(t, "Willi Buli", resultString)

	resultNumber := Length[int](100)
	assert.Equal(t, 100, resultNumber)
}

// ----------------------------------------------------------------------------------------------------------------------
func MultipleParameter[T1 any, T2 any](param1 T1, param2 T2) {
	fmt.Println(param1)
	fmt.Println(param2)
}

func TestMultipleParameter(t *testing.T) {
	MultipleParameter[string, int]("Willi Buli", 100)
	MultipleParameter[int, string](100, "Willi Buli")
}
