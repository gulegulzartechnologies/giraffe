package operations

import (
	"strings"

	"github.com/gulegulzartechnologies/giraffe"
)

// GroupBy groups the records on the basis of `columns`, and gathers `target` into a list of strings
func GroupBy(df *giraffe.Dataframe, columnsIndexes []int, targetIndex int) *giraffe.Dataframe {

	gb := &giraffe.Dataframe{
		ColumnNames: []string{df.ColumnNames[targetIndex]},
		Columns:     make(map[string]giraffe.Series),
		RowCount:    0,
	}

	for r := range df.RowCount {
		hashStringList := []string{}
		for i := range columnsIndexes {
			hashStringList = append(hashStringList, df.Columns[df.ColumnNames[i]].Values[r])
		}
		hashstring := strings.Join(hashStringList, "")
		hash := giraffe.Hasher(hashstring)

		if _, ok := gb.Columns[hash]; ok {
			series := gb.Columns[hash]
			series.Values = append(
				series.Values,
				df.Columns[df.ColumnNames[targetIndex]].Values[r],
			)
			gb.Columns[hash] = series
		} else {
			gb.Columns[hash] = giraffe.Series{
				Name:   hashstring,
				Values: []string{df.Columns[df.ColumnNames[targetIndex]].Values[r]},
			}
			gb.RowCount = gb.RowCount + 1
		}

	}

	return gb

}
