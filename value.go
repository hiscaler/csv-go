package csv

import (
	"errors"
	"strconv"
	"strings"
	"time"
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

func getValue(v Value, defaultValue ...string) string {
	s := v.String()
	if s == "" && len(defaultValue) > 0 {
		s = defaultValue[0]
	}
	return s
}

func (v Value) ToInt(defaultValue ...string) (int, error) {
	return strconv.Atoi(getValue(v, defaultValue...))
}

func (v Value) ToFloat64(defaultValue ...string) (float64, error) {
	return strconv.ParseFloat(getValue(v, defaultValue...), 64)
}

func (v Value) ToBool(defaultValue ...string) (bool, error) {
	return strconv.ParseBool(getValue(v, defaultValue...))
}

func (v Value) ToTime(layout string, loc *time.Location, defaultValue ...string) (time.Time, error) {
	s := getValue(v, defaultValue...)
	if s == "" {
		return time.Time{}, errors.New("is empty string")
	}
	return time.ParseInLocation(s, layout, loc)
}
