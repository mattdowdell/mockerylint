package usefactory

import (
	"testing"
)


func TestExample_Good(t *testing.T) {
	mock := NewMockExample(t)
	_ = mock
}

func TestExample_Bad(t *testing.T) {
	mock1 := new(MockExample) // want `use factory to initialise mock`
	_ = mock1

	mock2 := &MockExample{} // want `use factory to initialise mock`
	_ = mock2
}
