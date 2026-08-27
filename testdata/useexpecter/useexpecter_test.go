package useexpecter

import (
	"testing"
)


func TestExample_Good(t *testing.T) {
	mock := NewMockExample(t)
	mock.EXPECT().Example(1).Return(1).Once()

	got := mock.Example(1)

	if got != 1 {
		t.Fail()
	}
}

func TestExample_Bad(t *testing.T) {
	mock := NewMockExample(t)
	mock.On("Example", 1).Return(1).Once() // want `use .EXPECT instead of .On`

	got := mock.Example(1)

	if got != 1 {
		t.Fail()
	}
}

func TestExample_BadEmbedded(t *testing.T) {
	mock := NewMockExample(t)
	mock.Mock.On("Example", 1).Return(1).Once() // want `use .EXPECT instead of .On`

	got := mock.Example(1)

	if got != 1 {
		t.Fail()
	}
}

// TestLegacy_Good uses .On because the mock was generated without with-expecter, so it
// has no expecter to use instead.
func TestLegacy_Good(t *testing.T) {
	mock := NewMockLegacy(t)
	mock.On("Example", 1).Return(1).Once()

	got := mock.Example(1)

	if got != 1 {
		t.Fail()
	}
}

// TestLegacy_GoodEmbedded is TestLegacy_Good written against the embedded testify mock.
func TestLegacy_GoodEmbedded(t *testing.T) {
	mock := NewMockLegacy(t)
	mock.Mock.On("Example", 1).Return(1).Once()

	got := mock.Example(1)

	if got != 1 {
		t.Fail()
	}
}
