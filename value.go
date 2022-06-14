package csv

import (
	"strconv"
	"strings"
)

type Value struct {
	x             int
	y             int
	valid         bool
	OriginalValue string
	NewValue      string
}

func (v *Value) TrimSpace() *Value {
	v.NewValue = strings.TrimSpace(v.OriginalValue)
	return v
}

func (v *Value) Do(f func(s string) string) *Value {
	v.NewValue = f(v.OriginalValue)
	return v
}

func (v Value) String() string {
	return v.NewValue
}

func (v Value) ToInt(defaultValue ...string) (int, error) {
	s := v.String()
	if s == "" && len(defaultValue) > 0 {
		s = defaultValue[0]
	}
	return strconv.Atoi(s)
}

func (v Value) ToFloat64(defaultValue ...string) (float64, error) {
	s := v.String()
	if s == "" && len(defaultValue) > 0 {
		s = defaultValue[0]
	}
	return strconv.ParseFloat(s, 64)
}

func (v Value) ToBool(defaultValue ...string) (bool, error) {
	s := v.String()
	if s == "" && len(defaultValue) > 0 {
		s = defaultValue[0]
	}
	return strconv.ParseBool(s)
}
