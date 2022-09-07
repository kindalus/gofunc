package collections

type Collection[T any] []T

func From[T any](v ...T) Collection[T] {

	result := make([]T, len(v))
	copy(result, v)

	return Collection[T](result)
}

func (c Collection[T]) Filter(f func(r T) bool) Collection[T] {
	return Filter(c, f)
}

func (c Collection[T]) Any(f func(r T) bool) bool {
	return Any(c, f)
}
