package csv

import (
	"encoding/csv"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type CSV struct {
	file             *os.File    // Opened file
	reader           *csv.Reader // File reader
	currentRowNumber int         // Current row number
}

func NewCSV() *CSV {
	csv := &CSV{}
	defer func() {
		if csv.file != nil {
			csv.file.Close()
		}
	}()
	return csv
}

// Open opens a csv file
func (c *CSV) Open(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	c.file = f
	c.reader = csv.NewReader(f)
	if strings.EqualFold(filepath.Ext(filename), ".tsv") {
		c.reader.Comma = '\t'
	}
	c.currentRowNumber = 0
	return nil
}

func (c *CSV) Row() (r Row, isLastRow bool, err error) {
	record, err := c.reader.Read()
	if err != nil {
		isLastRow = err == io.EOF
		return
	}

	c.currentRowNumber += 1
	r = Row{
		valid:   true,
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

	csvWriter := csv.NewWriter(f)
	for i := range records {
		csvWriter.Write(records[i])
	}
	csvWriter.Flush()
	return csvWriter.Error()
}
