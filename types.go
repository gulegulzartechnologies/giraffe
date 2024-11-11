package giraffe

type Dataframe struct {
	ColumnNames []string
	Columns     map[string]Series
	RowCount    int
}

type Series struct {
	Name   string
	Values []string
}
