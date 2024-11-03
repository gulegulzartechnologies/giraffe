package types

import (
	"testing"
)

func TestNew(t *testing.T) {
	data := [][]string{
		{
			"String1",
			"String2",
		},
		{
			"String3",
			"String4",
		},
	}

	df := New(data)

	t.Log(df)
}
