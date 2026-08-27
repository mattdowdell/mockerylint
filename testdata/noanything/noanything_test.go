package noanything

import (
	"testing"

	"github.com/stretchr/testify/mock"
)


func TestExample_Good1(t *testing.T) {
	m := NewMockExample(t)

	m.EXPECT().Example(1, "a").Return(2).Once()
	m.EXPECT().Example(2, "b").RunAndReturn(func(int, string) int { return 4 }).Once()
}

// The alternatives to mock.Anything are all accepted.
func TestExample_Good2(t *testing.T) {
	m := NewMockExample(t)

	m.EXPECT().Example(mock.AnythingOfType("int"), "a").Return(2).Once()
	m.EXPECT().Example(1, mock.MatchedBy(func(s string) bool { return s != "" })).Return(4).Once()
	m.EXPECT().Example(mock.IsType(0), mock.AnythingOfType("string")).Return(6).Once()
}

func TestExample_Good3(t *testing.T) {
	m := NewMockExampleParam[int](t)

	m.EXPECT().Example(1).Return(2).Once()
	m.EXPECT().Example(mock.AnythingOfType("int")).Return(4).Once()
}

// Return values are not arguments, so a returned string that happens to hold the value of
// mock.Anything matches nothing and is left alone.
func TestExample_Good4(t *testing.T) {
	m := NewMockExample(t)

	m.EXPECT().Str().Return(mock.Anything).Once()
	m.EXPECT().Str().RunAndReturn(func() string { return mock.Anything }).Once()
}

// mock.Anything away from an expectation is not the concern of the rule.
func TestExample_Good5(t *testing.T) {
	m := NewMockExample(t)

	m.EXPECT().Example(1, "a").Return(2).Once()
	use(mock.Anything)

	if mock.Anything == "" {
		t.Fail()
	}
}

func use(_ string) {}

func TestExample_Bad1(t *testing.T) {
	m := NewMockExample(t)

	m.EXPECT().Example(mock.Anything, "a").Return(2).Once() // want `use the expected value, mock\.AnythingOfType, or mock\.MatchedBy instead of mock\.Anything`
	m.EXPECT().Example(1, mock.Anything).Return(4).Once() // want `use the expected value, mock\.AnythingOfType, or mock\.MatchedBy instead of mock\.Anything`
}

// Every offending argument of an expectation is reported.
func TestExample_Bad2(t *testing.T) {
	m := NewMockExample(t)

	m.EXPECT().Example(mock.Anything, mock.Anything).Return(2).Once() // want `use the expected value, mock\.AnythingOfType, or mock\.MatchedBy instead of mock\.Anything` `use the expected value, mock\.AnythingOfType, or mock\.MatchedBy instead of mock\.Anything`
}

func TestExample_Bad3(t *testing.T) {
	m := NewMockExampleParam[int](t)

	m.EXPECT().Example(mock.Anything).Return(2).Once() // want `use the expected value, mock\.AnythingOfType, or mock\.MatchedBy instead of mock\.Anything`
}

// The .On form matches arguments after the name of the mocked method.
func TestExample_Bad4(t *testing.T) {
	m := NewMockExample(t)

	m.On("Example", mock.Anything, "a").Return(2).Once() // want `use \.EXPECT instead of \.On` `use the expected value, mock\.AnythingOfType, or mock\.MatchedBy instead of mock\.Anything`
	m.Mock.On("Example", 1, mock.Anything).Return(4).Once() // want `use \.EXPECT instead of \.On` `use the expected value, mock\.AnythingOfType, or mock\.MatchedBy instead of mock\.Anything`
}

// Testify compares the argument by value, so an alias for mock.Anything or the string it
// holds written out in full matches anything just the same.
func TestExample_Bad5(t *testing.T) {
	const anything = mock.Anything

	m := NewMockExample(t)

	m.EXPECT().Example(1, anything).Return(2).Once() // want `use the expected value, mock\.AnythingOfType, or mock\.MatchedBy instead of mock\.Anything`
	m.EXPECT().Example(2, "mock.Anything").Return(4).Once() // want `use the expected value, mock\.AnythingOfType, or mock\.MatchedBy instead of mock\.Anything`
}

// An expectation that escapes is still judged on its arguments, as they are already given.
func TestExample_Bad6(t *testing.T) {
	m := NewMockExample(t)

	c := m.EXPECT().Example(mock.Anything, "a") // want `use the expected value, mock\.AnythingOfType, or mock\.MatchedBy instead of mock\.Anything`
	c.Return(2).Once()
}
