# mockerylint

## Rules

### Testify

The following rules are applied to mocks based on [testify/mock].

[testify/mock]: https://pkg.go.dev/github.com/stretchr/testify/mock

#### noassert

_In progress ([#4](https://github.com/mattdowdell/mockerylint/issues/4))._

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
func TestExample (t *testing.T) {
  example := NewExample(t)

  // add expectations here
  // use expectations here

  example.AssertExpectations()
}
```

</td><td>

```go
func TestExample (t *testing.T) {
  example := NewExample(t)

  // add expectations here
  // use expectations here

  example.AssertExpectations()
}
```

</td></tr>
</tbody></table>

#### usefactory

_In progress ([#5](https://github.com/mattdowdell/mockerylint/issues/5))._

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
func TestExample (t *testing.T) {
  example1 := new(Example)
  example2 := &Example{}
}
```

</td><td>

```go
func TestExample (t *testing.T) {
  example := NewExample(t)
}
```

</td></tr>
</tbody></table>

#### useexpecter

_In progress ([#3](https://github.com/mattdowdell/mockerylint/issues/3))._

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
func TestExample (t *testing.T) {
  example := NewExample(t)
  example.On("Thing").Return(nil).Once()
}
```

</td><td>

```go
func TestExample (t *testing.T) {
  example := NewExample(t)
  example.EXPECT().Thing().Return(nil).Once()
}
```

</td></tr>
</tbody></table>

#### usetimes

_In progress ([#6](https://github.com/mattdowdell/mockerylint/issues/6))._

_TODO: make allowing use of Maybe() configurable?_

<table>
<thead><tr><th>Bad</th><th>Good</th></tr></thead>
<tbody>
<tr><td>

```go
func TestExample (t *testing.T) {
  example := NewExample(t)
  example.On("Thing").Return(nil)
}
```

</td><td>

```go
func TestExample (t *testing.T) {
  example := NewExample(t)
  example.EXPECT().Thing1().Return(nil).Maybe()
  example.EXPECT().Thing2().Return(nil).Once()
  example.EXPECT().Thing3().Return(nil).Twice()
  example.EXPECT().Thing4().Return(nil).Times(3)
}
```

</td></tr>
</tbody></table>

## Matryer

Linting of mocks based on [matryer/moq] is not surrently supported.

[matryer/moq]: https://github.com/matryer/moq
