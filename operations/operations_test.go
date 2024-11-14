package operations

import (
	"testing"

	"github.com/gulegulzartechnologies/giraffe/io"
)

func TestGroupBy(t *testing.T) {

	testcsv := "../test-fixtures/Groceries_dataset.csv"

	df, err := io.ReadFromCSVWithHeadings(testcsv)
	if err != nil {
		t.Error(err)
	}

	t.Logf("df RowCount: %d", df.RowCount)

	gb := GroupBy(
		df,
		[]int{0, 1},
		2,
	)

	i := 0
	t.Log(gb.ColumnNames)
	for k, v := range gb.Columns {
		t.Log(k)
		t.Log(v.Name)
		t.Log(v.Values)
		i++
		if i == 5 {
			break
		}
	}
	t.Logf("gb RowCount: %d", gb.RowCount)

}

func TestHotEncode(t *testing.T) {
	tests := []struct {
		name            string
		path            string
		uniqueItemCount int
		expected        [][]int
	}{
		{
			name:            "simple hot encode",
			path:            "../test-fixtures/abstract.csv",
			uniqueItemCount: 40,
			expected:        [][]int{},
		},
		{
			name:            "simple hot encode",
			path:            "../test-fixtures/small.csv",
			uniqueItemCount: 7,
			expected:        [][]int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			df, err := io.ReadFromCSVWithHeadings(test.path)
			if err != nil {
				t.Error(err)
			}

			gb := GroupBy(
				df,
				[]int{0, 1},
				2,
			)

			he := HotEncode(gb, test.uniqueItemCount)

			t.Logf("items: %v", he.Mapping)

			for k, v := range he.Matrix {
				t.Logf("hash: %s \t enc: %v", k, v)
			}
		})
	}
}
