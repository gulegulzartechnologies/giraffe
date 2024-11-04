package giraffe

type Series[T comparable] struct {
	Name     string
	Values   T
	dataType string
}
