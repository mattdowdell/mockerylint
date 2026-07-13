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
