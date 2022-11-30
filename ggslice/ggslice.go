package ggslice

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func Map[From, To any](vs []From, f func(int, From) To) []To {
	result := make([]To, len(vs))
	for i, v := range vs {
		result[i] = f(i, v)
	}
	return result
}

func MapErr[From, To any](vs []From, f func(int, From) (To, error)) ([]To, error) {
	result := make([]To, len(vs))
	for i, v := range vs {
		t, err := f(i, v)
		if err != nil {
			return nil, err
		}
		result[i] = t
	}
	return result, nil
}

func Fold[From, To any](vs []From, initial To, f func(int, To, From) To) To {
	result := initial
	for i, t := range vs {
		result = f(i, result, t)
	}
	return result
}

func FoldErr[From, To any](vs []From, initial To, f func(int, To, From) (To, error)) (To, error) {
	result := initial
	for i, v := range vs {
		t, err := f(i, result, v)
		if err != nil {
			return nil, err
		}
		result = t
	}
	return result, nil
}

func Sort[T constraints.Ordered](ts []T) {
	sort.Slice(ts, func(i, j int) bool {
		return ts[i] < ts[j]
	})
}

func SortByField[T any, Field constraints.Ordered](ts []T, f func(T) Field) {
	sort.Slice(ts, func(i, j int) bool {
		return f(ts[i]) < f(ts[j])
	})
}
