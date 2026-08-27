package noanything

type Example interface {
	Example(int, string) int
	Str() string
}

type ExampleParam[T any] interface {
	Example(T) T
}
