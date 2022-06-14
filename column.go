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

func (c *Column) TrimSpace() *Column {
	c.NewValue = strings.TrimSpace(c.OriginalValue)
	return c
}

func (c *Column) Do(f func(s string) string) *Column {
	c.NewValue = f(c.OriginalValue)
	return c
}

func (c Column) String() string {
	return c.NewValue
}

func getValue(c Column, defaultValue ...string) string {
	s := c.String()
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

func (c Column) IsEmpty() bool {
	return c.NewValue == ""
}

func (c Column) IsBlack() bool {
	return c.NewValue == "" || strings.TrimSpace(c.NewValue) == ""
}

func (c Column) ToInt(defaultValue ...string) (int, error) {
	return strconv.Atoi(cleanNumber(getValue(c, defaultValue...)))
}

func (c Column) ToInt8(defaultValue ...string) (int8, error) {
	i, err := strconv.ParseInt(cleanNumber(getValue(c, defaultValue...)), 10, 8)
	if err != nil {
		return 0, err
	}
	return int8(i), nil
}

func (c Column) ToInt16(defaultValue ...string) (int16, error) {
	i, err := strconv.ParseInt(cleanNumber(getValue(c, defaultValue...)), 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(i), nil
}

func (c Column) ToInt32(defaultValue ...string) (int32, error) {
	i, err := strconv.ParseInt(cleanNumber(getValue(c, defaultValue...)), 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(i), nil
}

func (c Column) ToInt64(defaultValue ...string) (int64, error) {
	return strconv.ParseInt(cleanNumber(getValue(c, defaultValue...)), 10, 64)
}

func (c Column) ToFloat32(defaultValue ...string) (float32, error) {
	i, err := strconv.ParseFloat(cleanNumber(getValue(c, defaultValue...)), 32)
	if err != nil {
		return 0, err
	}
	return float32(i), nil
}

func (c Column) ToFloat64(defaultValue ...string) (float64, error) {
	return strconv.ParseFloat(cleanNumber(getValue(c, defaultValue...)), 64)
}

func (c Column) ToBool(defaultValue ...string) (bool, error) {
	return strconv.ParseBool(strings.ToLower(getValue(c, defaultValue...)))
}

func (c Column) ToTime(layout string, loc *time.Location, defaultValue ...string) (time.Time, error) {
	s := getValue(c, defaultValue...)
	if s == "" {
		return time.Time{}, errors.New("is empty string")
	}
	return time.ParseInLocation(s, layout, loc)
}
