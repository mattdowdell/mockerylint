package usefactory

import (
	"testing"

	"github.com/stretchr/testify/mock"
)


func TestExample_Good1(t *testing.T) {
	m := NewMockExample(t)
	_ = m
}

func TestExample_Good2(t *testing.T) {
	m := NewMockExampleParam[int](t)
	_ = m
}

// Declaring a pointer creates no mock, so it is left for the factory to fill in.
func TestExample_Good3(t *testing.T) {
	var m *MockExample
	m = NewMockExample(t)
	_ = m

	var ms []MockExample
	_ = ms
}

func TestExample_Ignore1(t *testing.T) {
	m := new(mock.Mock)
	m.Test(t)

	m.AssertExpectations(t)
}

func TestExample_Ignore2(t *testing.T) {
	m := &mock.Mock{}
	m.Test(t)

	m.AssertExpectations(t)
}

// A mock.Mock value is no more a generated mock than the pointers above.
func TestExample_Ignore3(t *testing.T) {
	var m mock.Mock
	m.Test(t)

	m.On("Example", 1).Return(1).Once()

	m.AssertExpectations(t)
}

// A mock.Mock reached through anything other than a generated mock is ignored too.
func TestExample_Ignore4(t *testing.T) {
	var w struct{ Mock mock.Mock }
	w.Mock.Test(t)

	w.Mock.AssertExpectations(t)
}

func identity(m *MockExample) *MockExample { return m }

// A shadowed new allocates nothing, so the factory is already in use here.
func TestExample_Ignore5(t *testing.T) {
	new := identity
	m := new(NewMockExample(t))
	_ = m
}

func TestExample_Bad1(t *testing.T) {
	m := new(MockExample) // want `use factory to initialise mock`
	m.Test(t) // want `\.Test\(\) can be removed when using mock factory`

	m.AssertExpectations(t) // want `\.AssertExpectations\(\) can be removed when using mock factory`
}

func TestExample_Bad2(t *testing.T) {
	m := &MockExample{} // want `use factory to initialise mock`
	m.Test(t) // want `\.Test\(\) can be removed when using mock factory`

	m.AssertExpectations(t) // want `\.AssertExpectations\(\) can be removed when using mock factory`
}

func TestExample_Bad3(t *testing.T) {
	m := new(MockExampleParam[int]) // want `use factory to initialise mock`
	m.Test(t) // want `\.Test\(\) can be removed when using mock factory`

	m.AssertExpectations(t) // want `\.AssertExpectations\(\) can be removed when using mock factory`
}

func TestExample_Bad4(t *testing.T) {
	m := &MockExampleParam[int]{} // want `use factory to initialise mock`
	m.Test(t) // want `\.Test\(\) can be removed when using mock factory`

	m.AssertExpectations(t) // want `\.AssertExpectations\(\) can be removed when using mock factory`
}

func TestExample_Bad5(t *testing.T) {
	m := NewMockExample(t)
	m.Test(t) // want `\.Test\(\) can be removed when using mock factory`

	m.AssertExpectations(t) // want `\.AssertExpectations\(\) can be removed when using mock factory`
}

func TestExample_Bad6(t *testing.T) {
	m := NewMockExample(t)
	m.Mock.Test(t) // want `\.Test\(\) can be removed when using mock factory`

	m.Mock.AssertExpectations(t) // want `\.AssertExpectations\(\) can be removed when using mock factory`
}

// The embedded mock is still reachable through a value receiver and through parentheses.
func TestExample_Bad7(t *testing.T) {
	var m MockExample // want `use factory to initialise mock`
	m.Mock.Test(t) // want `\.Test\(\) can be removed when using mock factory`

	(m.Mock).AssertExpectations(t) // want `\.AssertExpectations\(\) can be removed when using mock factory`
}

// A composite literal creates a mock whether or not it is addressed.
func TestExample_Bad8(t *testing.T) {
	m := MockExample{} // want `use factory to initialise mock`
	m.Test(t) // want `\.Test\(\) can be removed when using mock factory`

	n := MockExampleParam[int]{} // want `use factory to initialise mock`
	n.Test(t) // want `\.Test\(\) can be removed when using mock factory`
}

// A zero valued mock is usable, so it bypasses the factory too.
func TestExample_Bad9(t *testing.T) {
	var m MockExample // want `use factory to initialise mock`
	m.On("Example", 1).Return(1).Once() // want `use .EXPECT instead of .On`

	var n MockExampleParam[int] // want `use factory to initialise mock`
	n.Test(t) // want `\.Test\(\) can be removed when using mock factory`
}
