package golang_generics

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Data[T any] struct {
	First  T
	Second T
}

func TestData(t *testing.T) {
	data := Data[string]{
		First:  "Willi",
		Second: "Buli",
	}

	fmt.Println(data)
}

// ini untuk generic method ---------------------------------------------------------------------------------------------
func (d *Data[_]) SayHello(name string) string {
	return "Hallo " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return d.First
}

func TestGenericMethod(t *testing.T) {
	data := Data[string]{
		First:  "Willi",
		Second: "Buli",
	}

	assert.Equal(t, "Hallo Ama", data.SayHello("Ama"))
	assert.Equal(t, "Ama", data.ChangeFirst("Ama"))
}
