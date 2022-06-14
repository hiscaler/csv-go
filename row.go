package csv

type Row struct {
	valid  bool
	Number int
	Values []string
}

type Rows []Row

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

func (r *Row) Write(value *Value) *Row {
	if value.valid {
		r.Values[value.y] = value.String()
	}
	return r
}

func (r Row) Record() []string {
	return r.Values
}
