package csv

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

type Column struct {
	x             int
	y             int
	valid         bool
	OriginalValue string
	NewValue      string
}

func (v *Column) TrimSpace() *Column {
	v.NewValue = strings.TrimSpace(v.OriginalValue)
	return v
}

func (v *Column) Do(f func(s string) string) *Column {
	v.NewValue = f(v.OriginalValue)
	return v
}

func (v Column) String() string {
	return v.NewValue
}

func getValue(v Column, defaultValue ...string) string {
	s := v.String()
	if s == "" && len(defaultValue) > 0 {
		s = defaultValue[0]
	}
	return s
}

func (v Column) ToInt(defaultValue ...string) (int, error) {
	s := getValue(v, defaultValue...)
	if s != "" {
		s = strings.TrimSpace(s)
		s = strings.ReplaceAll(s, ",", "")
	}

	return strconv.Atoi(s)
}

func (v Column) ToFloat64(defaultValue ...string) (float64, error) {
	s := getValue(v, defaultValue...)
	if s != "" {
		s = strings.TrimSpace(s)
		s = strings.ReplaceAll(s, ",", "")
	}
	return strconv.ParseFloat(s, 64)
}

func (v Column) ToBool(defaultValue ...string) (bool, error) {
	return strconv.ParseBool(strings.ToLower(getValue(v, defaultValue...)))
}

func (v Column) ToTime(layout string, loc *time.Location, defaultValue ...string) (time.Time, error) {
	s := getValue(v, defaultValue...)
	if s == "" {
		return time.Time{}, errors.New("is empty string")
	}
	return time.ParseInLocation(s, layout, loc)
}
