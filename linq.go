package glinq

import "errors"

type Number interface {
	~int | ~int64 | ~float64 | ~float32 | ~uint | ~uint64
}

type Ordered interface {
	Number | string
}

var (
	ErrorCannotFound = errors.New("can not find element")
)
