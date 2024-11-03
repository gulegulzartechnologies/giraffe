package types

import (
	"reflect"
)

type metadata struct {
	dataType reflect.Type
}

type Dataframe[T comparable] struct {
	metadata metadata
	data     [][]T
}

func New[T comparable](data [][]T) Dataframe[T] {

	df := Dataframe[T]{
		metadata: metadata{
			dataType: reflect.TypeOf(data[0][0]),
		},
		data: data,
	}

	return df

}
