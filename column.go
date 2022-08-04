package csv

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

var numberReplacer *strings.Replacer

func init() {
	numberReplacer = strings.NewReplacer(",", "", " ", "")
}

// Column line of data in file
type Column struct {
	x             int    // Row number
	y             int    // Column number
	valid         bool   // Is valid column
	OriginalValue string // Original value
	NewValue      string // New value process after
}

// TrimSpace trim both space with original value
func (c *Column) TrimSpace() *Column {
	c.NewValue = strings.TrimSpace(c.NewValue)
	return c
}

func (c *Column) Do(f func(s string) string) *Column {
	c.NewValue = f(c.NewValue)
	return c
}

func (c Column) String() string {
	return c.NewValue
}

func getValue(c Column, defaultValue ...string) string {
	if c.NewValue != "" {
		return c.NewValue
	} else if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return ""
}

// Clean number string
// Reference: https://zhuanlan.zhihu.com/p/157980325
// Rules:
// 1,234 => 1234
// 123 456 => 123456
// Only support about two rules, if you have other rule, please use Do() method fixed the value.
func cleanNumber(s string) string {
	if s == "" {
		return ""
	}
	return numberReplacer.Replace(s)
}

func (c Column) IsEmpty() bool {
	return c.NewValue == ""
}

func (c Column) IsBlank() bool {
	return c.NewValue == "" || strings.TrimSpace(c.NewValue) == ""
}

func (c Column) IsNull() bool {
	return strings.EqualFold(c.NewValue, "NULL")
}

func (c Column) ToBytes(defaultValue ...string) []byte {
	s := getValue(c, defaultValue...)
	if s == "" {
		return []byte{}
	}
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
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
	return time.ParseInLocation(layout, s, loc)
}
