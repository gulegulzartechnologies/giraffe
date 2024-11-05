package giraffe

import (
	"testing"
)

func TestNew(t *testing.T) {

	testcsv := "./test-fixtures/iris.csv"

	df, err := ReadFromCSVWithHeadings(testcsv)
	if err != nil {
		t.Error(err)
	}

	t.Log(df.Columns)

}
