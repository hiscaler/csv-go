package csv

import (
	"encoding/csv"
	"errors"
	"github.com/dimchansky/utfbom"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type CSV struct {
	file             *os.File    // Opened file
	currentRowNumber int         // Current row number
	Reader           *csv.Reader // File Reader
}

func NewCSV() *CSV {
	return &CSV{}
}

// Get field delimiter from file extension name
func fieldDelimiter(ext string) rune {
	switch strings.ToLower(ext) {
	case ".tsv":
		return '\t'
	case ".psv":
		return '|'
	default:
		return ','
	}
}

// Open opens a csv file
func (c *CSV) Open(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	c.file = f
	c.Reader = csv.NewReader(utfbom.SkipOnly(f))
	c.Reader.FieldsPerRecord = 0
	c.Reader.Comma = fieldDelimiter(filepath.Ext(filename))
	c.currentRowNumber = 0
	return nil
}

// find Finds value by conditions and return row/column indexes
func (c *CSV) find(value string, fuzzy, all bool) (indexes []Index, err error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return indexes, errors.New("csv: find value is empty")
	}

	c.Reset()
	if fuzzy {
		value = strings.ToLower(value)
	}
	for {
		row, isEOF, e := c.Row()
		if isEOF {
			break
		}
		if e != nil {
			err = e
			return
		}

		maxColumns := len(row.Columns)
		for i := 1; i <= maxColumns; i++ {
			v := row.Column(i).TrimSpace().String()
			matched := false
			if fuzzy {
				matched = strings.Contains(strings.ToLower(v), value)
			} else {
				matched = strings.EqualFold(v, value)
			}
			if matched {
				index := Index{
					Row:    row.Number,
					Column: i,
				}
				if all {
					return []Index{index}, nil
				}
				indexes = append(indexes, index)
			}
		}
	}
	return
}

// FindAll Find all matched value row/column indexes
func (c *CSV) FindAll(value string, fuzzy bool) (indexes []Index, err error) {
	return c.find(value, fuzzy, false)
}

// FindFirst Find first matched value row/column index
func (c *CSV) FindFirst(value string, fuzzy bool) (index Index, err error) {
	indexes, err := c.find(value, fuzzy, true)
	if err == nil {
		if len(indexes) == 0 {
			err = errors.New("not found")
		} else {
			index = indexes[0]
		}
	}
	return
}

// FindLast Find last matched value row/column index
func (c *CSV) FindLast(value string, fuzzy bool) (index Index, err error) {
	indexes, err := c.find(value, fuzzy, false)
	if err == nil {
		n := len(indexes)
		if n == 0 {
			err = errors.New("not found")
		} else {
			index = indexes[n-1]
		}
	}
	return
}

// Close closes open file
func (c *CSV) Close() error {
	if c.file == nil {
		return nil
	}
	return c.file.Close()
}

// Reset resets to the file header, and set new Reader, used to re-read the file
func (c *CSV) Reset() error {
	if c.file == nil {
		return errors.New("file is closed")
	}
	_, err := c.file.Seek(0, 0)
	reader := csv.NewReader(utfbom.SkipOnly(c.file))
	reader.Comma = c.Reader.Comma
	reader.Comment = c.Reader.Comment
	reader.FieldsPerRecord = c.Reader.FieldsPerRecord
	reader.LazyQuotes = c.Reader.LazyQuotes
	reader.TrimLeadingSpace = c.Reader.TrimLeadingSpace
	reader.ReuseRecord = c.Reader.ReuseRecord
	c.Reader = reader
	c.currentRowNumber = 0
	return err
}

// Row read one row from opened file
func (c *CSV) Row() (r Row, isEOF bool, err error) {
	record, err := c.Reader.Read()
	if err != nil {
		if err == io.EOF {
			isEOF = true
			err = nil
		}
		return
	}

	c.currentRowNumber += 1
	r = Row{
		Number:  c.currentRowNumber,
		Columns: record,
	}
	return
}

// SaveAs save as file
func (c *CSV) SaveAs(filename string, records [][]string) error {
	dir := filepath.Dir(filename)
	_, err := os.Stat(dir)
	if err != nil && !os.IsExist(err) {
		// Creates dir
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	writer.Comma = fieldDelimiter(filepath.Ext(filename))
	err = writer.WriteAll(records)
	if err != nil {
		return err
	}
	return writer.Error()
}
