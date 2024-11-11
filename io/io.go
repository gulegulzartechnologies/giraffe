package io

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/gulegulzartechnologies/giraffe"
)

func ReadFromCSVWithHeadings(path string) (*giraffe.Dataframe, error) {

	df := &giraffe.Dataframe{
		ColumnNames: []string{},
		Columns:     make(map[string]giraffe.Series),
		RowCount:    0,
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV file: %v", err)
	}

	df.ColumnNames = append(df.ColumnNames, records[0]...)

	for _, row := range records[1:] {
		for i, col := range df.ColumnNames {
			series := df.Columns[col]
			series.Values = append(series.Values, row[i])
			df.Columns[col] = series
		}
		df.RowCount = df.RowCount + 1
	}

	return df, nil

}
