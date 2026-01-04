package usefactory

import (
	"testing"
)


func TestExample_Good1(t *testing.T) {
	mock := NewMockExample(t)
	_ = mock
}

func TestExample_Good2(t *testing.T) {
	mock := NewMockExampleParam[int](t)
	_ = mock
}

func TestExample_Bad1(t *testing.T) {
	mock := new(MockExample) // want `use factory to initialise mock`
	mock.Test(t) // want `\.Test\(\) can be removed when using mock factory`

	mock.AssertExpectations(t) // want `\.AssertExpectations\(\) can be removed when using mock factory`
}

func TestExample_Bad2(t *testing.T) {
	mock := &MockExample{} // want `use factory to initialise mock`
	mock.Test(t) // want `\.Test\(\) can be removed when using mock factory`

	mock.AssertExpectations(t) // want `\.AssertExpectations\(\) can be removed when using mock factory`
}

func TestExample_Bad3(t *testing.T) {
	mock := new(MockExampleParam[int]) // want `use factory to initialise mock`
	mock.Test(t) // want `\.Test\(\) can be removed when using mock factory`

	mock.AssertExpectations(t) // want `\.AssertExpectations\(\) can be removed when using mock factory`
}

func TestExample_Bad4(t *testing.T) {
	mock := &MockExampleParam[int]{} // want `use factory to initialise mock`
	mock.Test(t) // want `\.Test\(\) can be removed when using mock factory`

	mock.AssertExpectations(t) // want `\.AssertExpectations\(\) can be removed when using mock factory`
}

func TestExample_Bad5(t *testing.T) {
	mock := NewMockExample(t)
	mock.Test(t) // want `\.Test\(\) can be removed when using mock factory`

	mock.AssertExpectations(t) // want `\.AssertExpectations\(\) can be removed when using mock factory`
}

func TestExample_Bad6(t *testing.T) {
	mock := NewMockExample(t)
	mock.Mock.Test(t) // want `\.Test\(\) can be removed when using mock factory`

	mock.Mock.AssertExpectations(t) // want `\.AssertExpectations\(\) can be removed when using mock factory`
}
