package usetimes

import (
	"testing"
)


func TestExample_Good1(t *testing.T) {
	m := NewMockExample(t)

	m.EXPECT().Example(1).Return(2).Maybe()
	m.EXPECT().Example(2).Return(4).Once()
	m.EXPECT().Example(3).Return(6).Twice()
	m.EXPECT().Example(4).Return(8).Times(3)
}

func TestExample_Good2(t *testing.T) {
	m := NewMockExample(t)

	m.EXPECT().Example(1).RunAndReturn(func(int) int { return 2 }).Maybe()
	m.EXPECT().Example(2).RunAndReturn(func(int) int { return 4 }).Once()
	m.EXPECT().Example(3).RunAndReturn(func(int) int { return 6 }).Twice()
	m.EXPECT().Example(4).RunAndReturn(func(int) int { return 8 }).Times(3)
}

func TestExample_Good3(t *testing.T) {
	m := NewMockExample(t)

	m.EXPECT().Example(1).Run(func(int) {}).Return(2).Maybe()
	m.EXPECT().Example(2).Run(func(int) {}).Return(4).Once()
	m.EXPECT().Example(3).Run(func(int) {}).Return(6).Twice()
	m.EXPECT().Example(4).Run(func(int) {}).Return(8).Times(3)
}

func TestExample_Good4(t *testing.T) {
	m := NewMockExampleParam[int](t)

	m.EXPECT().Example(1).Return(2).Maybe()
	m.EXPECT().Example(2).Return(4).Once()
	m.EXPECT().Example(3).Return(6).Twice()
	m.EXPECT().Example(4).Return(8).Times(3)
}

func TestExample_Good5(t *testing.T) {
	m := NewMockExampleParam[int](t)

	m.EXPECT().Example(1).RunAndReturn(func(int) int { return 2 }).Maybe()
	m.EXPECT().Example(2).RunAndReturn(func(int) int { return 4 }).Once()
	m.EXPECT().Example(3).RunAndReturn(func(int) int { return 6 }).Twice()
	m.EXPECT().Example(4).RunAndReturn(func(int) int { return 8 }).Times(3)
}

func TestExample_Good6(t *testing.T) {
	m := NewMockExampleParam[int](t)

	m.EXPECT().Example(1).Run(func(int) {}).Return(2).Maybe()
	m.EXPECT().Example(2).Run(func(int) {}).Return(4).Once()
	m.EXPECT().Example(3).Run(func(int) {}).Return(6).Twice()
	m.EXPECT().Example(4).Run(func(int) {}).Return(8).Times(3)
}

func TestExample_Bad1(t *testing.T) {
	m := NewMockExample(t)

	m.EXPECT().Example(1).Return(2) // want `expectation should call \.Maybe\(\), \.Once\(\), \.Twice\(\), or \.Times\(N\)`
	m.EXPECT().Example(2).Run(func(_ int) {}).Return(4) // want `expectation should call \.Maybe\(\), \.Once\(\), \.Twice\(\), or \.Times\(N\)`
	m.EXPECT().Example(3).RunAndReturn(func(_ int) int { return 6 }) // want `expectation should call \.Maybe\(\), \.Once\(\), \.Twice\(\), or \.Times\(N\)`
}

func TestExample_Bad2(t *testing.T) {
	m := NewMockExampleParam[int](t)

	m.EXPECT().Example(1).Return(2) // want `expectation should call \.Maybe\(\), \.Once\(\), \.Twice\(\), or \.Times\(N\)`
	m.EXPECT().Example(2).Run(func(_ int) {}).Return(4) // want `expectation should call \.Maybe\(\), \.Once\(\), \.Twice\(\), or \.Times\(N\)`
	m.EXPECT().Example(3).RunAndReturn(func(_ int) int { return 6 }) // want `expectation should call \.Maybe\(\), \.Once\(\), \.Twice\(\), or \.Times\(N\)`
}

// A times method later in the chain satisfies the rule.
func TestExample_Good7(t *testing.T) {
	m := NewMockExample(t)

	m.EXPECT().Example(1).Return(2).Run(func(_ int) {}).Once()
	m.EXPECT().Example(2).Return(4).Panic("boom").Maybe()
	(m.EXPECT().Example(3).Return(6)).Once()
}

// An expectation that escapes may have a times method called on it elsewhere, so it
// cannot be judged in isolation.
func TestExample_Good8(t *testing.T) {
	m := NewMockExample(t)

	c := m.EXPECT().Example(1).Return(2)
	c.Once()

	d := m.EXPECT().Example(2).Return(4).Run(func(_ int) {})
	d.Twice()

	use(m.EXPECT().Example(3).Return(6))
	_ = build(m)
}

func use(_ *MockExample_Example_Call) {}

func build(m *MockExample) *MockExample_Example_Call {
	return m.EXPECT().Example(4).Return(8)
}

// A chain that ends in a non-times method is still missing a times method.
func TestExample_Bad3(t *testing.T) {
	m := NewMockExample(t)

	m.EXPECT().Example(1).Return(2).Run(func(_ int) {}) // want `expectation should call \.Maybe\(\), \.Once\(\), \.Twice\(\), or \.Times\(N\)`
	m.EXPECT().Example(2).Return(4).Panic("boom") // want `expectation should call \.Maybe\(\), \.Once\(\), \.Twice\(\), or \.Times\(N\)`
	m.EXPECT().Example(3).RunAndReturn(func(_ int) int { return 6 }).Run(func(_ int) {}) // want `expectation should call \.Maybe\(\), \.Once\(\), \.Twice\(\), or \.Times\(N\)`
}

// Expectations nested inside a chained closure are judged on their own.
func TestExample_Bad4(t *testing.T) {
	m := NewMockExample(t)
	other := NewMockExample(t)

	m.EXPECT().Example(1).Run(func(_ int) {
		other.EXPECT().Example(2).Return(4) // want `expectation should call \.Maybe\(\), \.Once\(\), \.Twice\(\), or \.Times\(N\)`
	}).Return(2).Once()
}
