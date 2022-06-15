package csv

import "strings"

type Row struct {
	Number  int
	Columns []string
}

type Rows []Row

func (r Row) IsEmpty() bool {
	return strings.Join(r.Columns, "") == ""
}

// Column reads column value in current row
func (r Row) Column(index int) *Column {
	c := &Column{
		x:             r.Number,
		y:             index,
		valid:         false,
		OriginalValue: "",
		NewValue:      "",
	}
	if index < len(r.Columns) {
		value := r.Columns[index]
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
func (r Row) Every(f func(r Row) bool) bool {
	return f(r)
}

// Map process all columns value in current row
func (r *Row) Map(f func(s string) string, columnIndex ...int) *Row {
	all := len(columnIndex) == 0
	for i, s := range r.Columns {
		if !all {
			next := false
			for _, j := range columnIndex {
				if i == j {
					next = true
					break
				}
			}
			if !next {
				continue
			}
		}
		r.Columns[i] = f(s)
	}
	return r
}

// Record change to string slice
func (r Row) Record() []string {
	return r.Columns
}
