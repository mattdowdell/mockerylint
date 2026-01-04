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
