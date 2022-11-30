package ggonce

import (
	"sync"

	"github.com/morikuni/go-generics/ggopt"
)

type Once[T any] struct {
	value ggopt.Option[T]
	mu    sync.RWMutex
}

func (o *Once[T]) tryGet() (T, bool) {
	o.mu.RLock()
	defer o.mu.RUnlock()

	v, ok := ggopt.Get(o.value)

	return v, ok
}

func (o *Once[T]) Do(f func() (T, error)) (T, error) {
	if v, ok := ggopt.Get(o.value); ok {
		return v, nil
	}
	o.mu.Lock()
	defer o.mu.Unlock()

	if v, ok := ggopt.Get(o.value); ok {
		return v, nil
	}

	v, err := f()
	if err != nil {
		// Return v because creating zero value of T is annoying.
		return v, err
	}

	o.value = ggopt.Some(v)

	return v, nil
}

func (o *Once[T]) MustDo(f func() (T, error)) T {
	v, err := o.Do(f)
	if err != nil {
		panic(err)
	}

	return v
}
