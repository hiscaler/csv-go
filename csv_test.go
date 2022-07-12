package csv

import (
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"strings"
	"testing"
)

var csvInstance *CSV

func TestMain(m *testing.M) {
	csvInstance = NewCSV()
	err := csvInstance.Open("./testdata/test.csv")
	if err != nil {
		panic(err)
	}
	defer csvInstance.Close()
	m.Run()
}

func TestCSV(t *testing.T) {
	for {
		row, isEOF, err := csvInstance.Row()
		if isEOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// Column first column data with current row, and add "A" prefix return column
		column := row.Column(0).
			TrimSpace().
			Do(func(s string) string {
				return s + "A"
			})
		switch row.Number {
		case 2:
			assert.Equal(t, "oneA", column.String())
		case 3:
			assert.Equal(t, "twoA", column.String())
		case 4:
			assert.Equal(t, "threeA", column.String())
		}

		column = row.Column(0).
			TrimSpace().
			Do(func(s string) string {
				// change return column
				v := ""
				switch s {
				case "one":
					v = "1"
				case "two":
					v = "2"
				case "three":
					v = "3"
				case "four":
					v = "4"
				case "five":
					v = "5"
				default:
					v = ""
				}
				return v
			})
		if row.Number != 1 {
			i, _ := column.ToInt()
			assert.Equal(t, row.Number-1, i)
		}
	}
}

func TestTSV(t *testing.T) {
	tsvInstance := NewCSV()
	err := tsvInstance.Open("./testdata/test.csv")
	if err != nil {
		panic(err)
	}
	defer tsvInstance.Close()
	for {
		row, isEOF, err := csvInstance.Row()
		if isEOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// Column first column data with current row, and add "A" prefix return value
		column := row.Column(0).
			TrimSpace().
			Do(func(s string) string {
				return s + "A"
			})
		switch row.Number {
		case 2:
			assert.Equal(t, "oneA", column.String())
		case 3:
			assert.Equal(t, "twoA", column.String())
		case 4:
			assert.Equal(t, "threeA", column.String())
		}

		column = row.Column(0).
			TrimSpace().
			Do(func(s string) string {
				// change return value
				v := ""
				switch s {
				case "one":
					v = "1"
				case "two":
					v = "2"
				case "three":
					v = "3"
				default:
					v = ""
				}
				return v
			})
		if row.Number != 1 {
			i, _ := column.ToInt()
			assert.Equal(t, row.Number-1, i)
		}
	}
}

func TestRowMap(t *testing.T) {
	const prefix = "PREFIX_"
	for {
		row, isEOF, err := csvInstance.Row()
		if isEOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		row.Map(func(s string) string {
			return prefix + s
		}, 0)
		row.Map(func(s string) string {
			return prefix + s
		}, 1)
		assert.Equal(t, true, strings.HasPrefix(row.Columns[0], prefix), "row.map.column 0")
		assert.Equal(t, true, strings.HasPrefix(row.Columns[1], prefix), "row.map.column 1")
		assert.Equal(t, false, strings.HasPrefix(row.Columns[2], prefix), "row.map.column 2")
	}
}

func TestRowEvery(t *testing.T) {
	exists := false
	hasNullValue := false
	err := csvInstance.Reset()
	assert.Equal(t, nil, err, "csv.reset method")

	for {
		row, isEOF, err := csvInstance.Row()
		if isEOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		exists = row.Every(func(r Row) bool {
			age, e := r.Column(2).ToInt()
			if e == nil && age > 20 {
				exists = true
				return exists
			}
			return false
		})
		hasNullValue = row.Every(func(r Row) bool {
			return r.Column(1).IsNull()
		})
	}
	assert.Equal(t, true, exists, "row.every")
	assert.Equal(t, true, hasNullValue, "row.every.null-check")
}

func TestSaveAs(t *testing.T) {
	records := make([][]string, 0)
	for {
		row, isEOF, err := csvInstance.Row()
		if isEOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		row.Map(func(s string) string {
			if row.Number != 1 {
				// Ignore header
				s = `1, "change"` + s
			}
			return s
		})
		records = append(records, row.Columns)
	}
	err := csvInstance.SaveAs("./testdata/a/a.csv", records)
	assert.Equal(t, nil, err, "save as")
	if err == nil {
		err = os.RemoveAll("./testdata/a")
		assert.Equal(t, nil, err, "remove dir")
	}
}
