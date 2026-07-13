package usetimes

type Example interface {
	Example(int) int
}

type ExampleParam[T any] interface {
	Example(T) T
}
