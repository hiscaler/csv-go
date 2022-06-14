CSV For Golang
==============

csv-go is csv file helper function, support csv, tsv format

## Install

```go
go get github.com/hiscaler/csv-go
```

## Notices

1. **Row start index is 1**
2. **Column start index is 0**

## Usage

### Open file
```go
csv := NewCSV()
err := csv.Open("./testdata/test.csv", 1, 2)
if err != nil {
    t.Error(err)
}
```

### Read
```go
for {
    row, isLastRow, err := csv.Row()
    if isLastRow {
        break
    }
    if err != nil {
        log.Fatal(err)
    }
    // Read first column data with current row, and add "A" prefix return value
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

    column = row.Read(0).
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
```

### Read a row
```go
row, isLastRow, err := csv.Row()
```

### Change row data

If you want to fix column value in current row, you can do it：

```go
row.Map(func (s string) string {
    return "PREFIX_" + s
}, 0)
```

The above code change first column value, will return "PREFIX_" and original column value concatenated string.

If you want change all columns value, don't pass `columnIndex` parameter value. Then all columns value will add "PREFIX_" prefix string.


### Read column in current row
```go
// Read first column in current row
column := row.Column(0)
column.TrimSpace() // Clear spaces
column.Do(func(s string) string {
	// process value
	return s
}) // Use Do method process value
column.String() // get value with string

// if you want to get correct value, you will check err and continue
v, err := column.ToInt() // get int value
v, err := column.ToInt("100") // get int value, and return 100 if value is empty
v, err := column.ToFloat64() // get float value
v, err := column.ToFloat64("100.00") // get float value, and return 100.00 if value is empty
v, err := column.ToBool() // get boolean value
v, err := column.ToBool("false") // get boolean value, and return false if value is empty
v, err := column.ToTime() // get time value
v, err := column.ToTime("2022-01-01") // get time value, and return 2022-01-01 if value is empty
```