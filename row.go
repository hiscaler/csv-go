package csv

type Row struct {
	valid  bool
	Number int
	Values []string
}

type Rows []Row

// Read reads column value in current row
func (r Row) Read(index int) *Value {
	valid := false
	value := ""
	if index < len(r.Values) {
		value = r.Values[index]
		valid = true
	}
	return &Value{
		x:             r.Number,
		y:             index,
		valid:         valid,
		OriginalValue: value,
		NewValue:      value,
	}
}

// Write writes column value in current row
func (r *Row) Write(value *Value) *Row {
	if value.valid {
		r.Values[value.y] = value.String()
	}
	return r
}

func (r Row) Every(f func(r Row) bool) bool {
	return f(r)
}

// Record change to string slice
func (r Row) Record() []string {
	return r.Values
}
