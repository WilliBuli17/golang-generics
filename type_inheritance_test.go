package golang_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Employee interface {
	GetName() string
}

func GetName[T Employee](param T) string {
	return param.GetName()
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

type VicePresident interface {
	GetName() string
	GetVicePresident() string
}

type MyVicePresident struct {
	Name string
}

func (m *MyVicePresident) GetName() string {
	return m.Name
}

func (m *MyVicePresident) GetVicePresident() string {
	return m.Name
}

func TestGetName(t *testing.T) {
	assert.Equal(t, "Willi", GetName[Manager](&MyManager{Name: "Willi"}))
	assert.Equal(t, "Buli", GetName[VicePresident](&MyVicePresident{Name: "Buli"}))
}
