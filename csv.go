package csv

import (
	"encoding/csv"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type CSV struct {
	file               *os.File
	reader             *csv.Reader
	headerRowNumber    int         // Header row number
	dataStartRowNumber int         // Data start row number
	currentRowNumber   int         // Current row number
	rows               map[int]Row // All data rows
	HeaderRow          Header      // Header row
	DataRows           Rows        // Data rows
}

func NewCSV() *CSV {
	csv := &CSV{
		rows: make(map[int]Row, 0),
	}
	defer csv.file.Close()
	return csv
}

// Open opens a csv file
func (c *CSV) Open(filename string, headerRowNumber, dataStartRowNumber int) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	c.file = f
	c.reader = csv.NewReader(f)
	if strings.EqualFold(filepath.Ext(filename), ".tsv") {
		c.reader.Comma = '\t'
	}
	c.headerRowNumber = headerRowNumber
	c.dataStartRowNumber = dataStartRowNumber
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
	if c.currentRowNumber == c.headerRowNumber {
		c.HeaderRow = record
	} else {
		c.DataRows = append(c.DataRows, r)
	}
	c.rows[c.currentRowNumber] = r
	return
}
