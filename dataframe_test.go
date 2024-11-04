package giraffe

import (
	"testing"
)

func TestNew(t *testing.T) {
	_ = [][]string{
		{
			"String1",
			"String2",
		},
		{
			"String3",
			"String4",
		},
	}

	df := New()

	t.Log(df)
}
