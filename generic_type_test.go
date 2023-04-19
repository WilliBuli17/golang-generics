package golang_generics

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, value := range bag {
		fmt.Println(value)
	}
}

func TestPrintBag(t *testing.T) {
	names := Bag[string]{"Ama", "Willi", "Buli"}
	PrintBag(names)

	numbers := Bag[int]{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	PrintBag(numbers)
}
