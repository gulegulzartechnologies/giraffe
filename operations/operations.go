package operations

import (
	"log/slog"
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

// HotEncode creates a one-hot encoded matrix.
// Accepts a transactional dataframe and the count of unique items.
// Returns a 2D int slice
func HotEncode(df *giraffe.Dataframe, itemCount int) *giraffe.HotEncoded {

	enc := &giraffe.HotEncoded{
		Mapping:   make(map[string]int),
		Matrix:    make(map[string][]int),
		ItemCount: itemCount,
	}

	for hash, col := range df.Columns {

		// initialize row encoding as 0
		enc.Matrix[hash] = make([]int, enc.ItemCount)
		for i := range enc.Matrix[hash] {
			enc.Matrix[hash][i] = 0
		}
		slog.Info("[HOTENCODE]", "hash", hash, "column", col)

		// update encoding for items
		for _, val := range col.Values {
			slog.Info("[HOTENCODE]", "val", val)
			if _, ok := enc.Mapping[val]; !ok {
				enc.Mapping[val] = len(enc.Mapping)
			}
			enc.Matrix[hash][enc.Mapping[val]] = 1
		}
	}

	return enc
}
