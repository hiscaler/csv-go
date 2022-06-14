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

func cleanNumber(s string) string {
	if s != "" {
		s = strings.TrimSpace(s)
		s = strings.ReplaceAll(s, ",", "")
	}
	return s
}

func (v Column) ToInt(defaultValue ...string) (int, error) {
	return strconv.Atoi(cleanNumber(getValue(v, defaultValue...)))
}

func (v Column) ToInt8(defaultValue ...string) (int8, error) {
	i, err := strconv.ParseInt(cleanNumber(getValue(v, defaultValue...)), 10, 8)
	if err != nil {
		return 0, err
	}
	return int8(i), nil
}

func (v Column) ToInt16(defaultValue ...string) (int16, error) {
	i, err := strconv.ParseInt(cleanNumber(getValue(v, defaultValue...)), 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(i), nil
}

func (v Column) ToInt32(defaultValue ...string) (int32, error) {
	i, err := strconv.ParseInt(cleanNumber(getValue(v, defaultValue...)), 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(i), nil
}

func (v Column) ToInt64(defaultValue ...string) (int64, error) {
	return strconv.ParseInt(cleanNumber(getValue(v, defaultValue...)), 10, 64)
}

func (v Column) ToFloat32(defaultValue ...string) (float32, error) {
	i, err := strconv.ParseFloat(cleanNumber(getValue(v, defaultValue...)), 32)
	if err != nil {
		return 0, err
	}
	return float32(i), nil
}

func (v Column) ToFloat64(defaultValue ...string) (float64, error) {
	return strconv.ParseFloat(cleanNumber(getValue(v, defaultValue...)), 64)
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
