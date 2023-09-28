package csv

type Value struct {
	Row  int
	Col  int
	Data []string
}

func (v *Value) Column(i int) *Column {
	c := &Column{
		x:     v.Row,
		y:     v.Col,
		valid: false,
	}
	if i <= 0 || i > len(v.Data) {
		return c
	}

	value := v.Data[i-1]
	c.valid = true
	c.OriginalValue = value
	c.NewValue = value

	return c
}
