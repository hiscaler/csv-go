package csv

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var csvInstance *CSV

func TestMain(m *testing.M) {
	csvInstance = NewCSV()
	err := csvInstance.Open("./testdata/test.csv", 1, 2)
	if err != nil {
		panic(err)
	}
	m.Run()
}

func TestCSV(t *testing.T) {
	for {
		row, isLastRow, err := csvInstance.Row()
		if isLastRow {
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
	err := tsvInstance.Open("./testdata/test.csv", 1, 2)
	if err != nil {
		panic(err)
	}
	for {
		row, isLastRow, err := csvInstance.Row()
		if isLastRow {
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
	for {
		row, isLastRow, err := csvInstance.Row()
		if isLastRow {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		row.Map(func(s string) string {
			return "PREFIX_" + s
		}, 0)
		row.Map(func(s string) string {
			return "PREFIX_" + s
		}, 1)
	}
	assert.Equal(t, "PREFIX_one", csvInstance.rows[2].Columns[0], "row.map")
	assert.Equal(t, "PREFIX_two", csvInstance.rows[3].Columns[0], "row.map")
	assert.Equal(t, "PREFIX_three", csvInstance.rows[4].Columns[0], "row.map")
	assert.Equal(t, "PREFIX_A", csvInstance.rows[2].Columns[1], "row.map")
	assert.Equal(t, "10", csvInstance.rows[2].Columns[2], "row.map")
}

func TestRowEvery(t *testing.T) {
	exists := false
	for {
		row, isLastRow, err := csvInstance.Row()
		if isLastRow {
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
		if exists {
			break
		}
	}
	assert.Equal(t, true, exists, "row.every")
}
