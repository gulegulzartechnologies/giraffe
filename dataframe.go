package giraffe

type DataFrame[T comparable] struct {
	// ColumnNames []string
	Columns  map[string]Series[T]
	RowCount int
}
