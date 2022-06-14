package csv

import (
	"encoding/csv"
	"os"
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
	c.headerRowNumber = headerRowNumber
	c.dataStartRowNumber = dataStartRowNumber
	c.currentRowNumber = 0
	return nil
}

func (c *CSV) Read() (r Row, err error) {
	record, err := c.reader.Read()
	if err != nil {
		return
	}

	c.currentRowNumber += 1
	r = Row{
		valid:  true,
		Number: c.currentRowNumber,
		Values: record,
	}
	if c.currentRowNumber == c.headerRowNumber {
		c.HeaderRow = record
	} else {
		c.DataRows = append(c.DataRows, r)
	}
	c.rows[c.currentRowNumber] = r
	return
}
