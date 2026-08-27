package useexpecter

type Example interface {
	Example(int) int
}

// Legacy is mocked without with-expecter enabled, as Mockery allowed before v3.0.0.
type Legacy interface {
	Example(int) int
}
