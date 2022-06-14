package csv

import (
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"testing"
)

func TestNewCSV(t *testing.T) {
	csv := NewCSV()
	err := csv.Open("./testdata/test.csv", 1, 2)
	if err != nil {
		t.Error(err)
	}

	for {
		row, err := csv.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// Read first column data with current row, and add "A" prefix return value
		value := row.Read(0).
			TrimSpace().
			Do(func(s string) string {
				return s + "A"
			})
		switch row.Number {
		case 2:
			assert.Equal(t, "oneA", value.String())
		case 3:
			assert.Equal(t, "twoA", value.String())
		case 4:
			assert.Equal(t, "threeA", value.String())
		}

		value = row.Read(0).
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
			i, _ := value.ToInt()
			assert.Equal(t, row.Number-1, i)
		}
	}
}
