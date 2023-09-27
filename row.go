package csv

import (
	"sort"
	"strings"
)

type Row struct {
	Number  int
	Columns []string
}

type Rows []Row

func (r *Row) IsEmpty() bool {
	return strings.Join(r.Columns, "") == ""
}

// Column reads column value in current row, column index start 1
func (r *Row) Column(index int) *Column {
	c := &Column{
		x:             r.Number,
		y:             index,
		valid:         false,
		OriginalValue: "",
		NewValue:      "",
	}
	if index > 0 && index <= len(r.Columns) {
		value := r.Columns[index-1]
		c.OriginalValue = value
		c.NewValue = value
		c.valid = true
	}
	return c
}

// Write writes column value in current row
func (r *Row) Write(column *Column) *Row {
	if column.valid {
		r.Columns[column.y] = column.String()
	}
	return r
}

// Every check condition is passed for all columns value in current row
func (r *Row) Every(f func(r *Row) bool) bool {
	return f(r)
}

// Map process all columns value in current row
func (r *Row) Map(f func(s string) string, columnIndex ...int) *Row {
	n := len(columnIndex)
	if n > 1 && !sort.IntsAreSorted(columnIndex) {
		sort.Ints(columnIndex)
	}
	for i, s := range r.Columns {
		do := false
		switch n {
		case 0:
			do = true
		case 1:
			do = i+1 == columnIndex[0]
		default:
			index := sort.SearchInts(columnIndex, i+1)
			do = index < n && columnIndex[index] == i+1
		}
		if do {
			r.Columns[i] = f(s)
		}
	}
	return r
}
