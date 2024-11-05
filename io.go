package giraffe

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadFromCSVWithHeadings(path string) (*DataFrame[string], error) {

	df := &DataFrame[string]{
		Columns: make(map[string]Series[string]),
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

	for _, row := range records[1:] {
		for i, col := range records[0] {
			series := df.Columns[col]
			series.Values = append(series.Values, row[i])
			df.Columns[col] = series
		}
	}

	return df, nil

}
