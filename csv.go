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
	csv := &CSV{}
	return csv
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
	c.Reader.FieldsPerRecord = -1
	c.Reader.Comma = fieldDelimiter(filepath.Ext(filename))
	c.currentRowNumber = 0
	return nil
}

// Close closes open file
func (c CSV) Close() error {
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
func (c CSV) SaveAs(filename string, records [][]string) error {
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
