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
