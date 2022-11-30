package ggopt

type Option[T any] *T

func Some[T any](t T) Option[T] {
	return &t
}

func Get[T any](o Option[T]) (zero T, exists bool) {
	if o == nil {
		return zero, false
	}

	return *o, true
}

func GetOrElse[T any](o Option[T], or T) T {
	if t, ok := Get(o); ok {
		return t
	}
	return or
}

func Map[T, U any](o Option[T], f func(t T) U) Option[U] {
	if t, ok := Get(o); ok {
		return Some(f(t))
	}
	return nil
}
