package usefactory

type Example interface {
	Example(int) int
}

type ExampleParam[T any] interface {
	Example(T) T
}

// Legacy is mocked without a factory, as Mockery generated before v2.11.0.
type Legacy interface {
	Example(int) int
}
