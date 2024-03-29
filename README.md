CSV For Golang
==============

![CI status](https://github.com/hiscaler/csv-go/actions/workflows/ci.yml/badge.svg)

csv-go is csv/tsv/psv file helper, current supported csv, tsv, psv format.

Use it to help you process data quickly.

## Install

```go
go get github.com/hiscaler/csv-go
```

## Notices

**Row and Column start index from 1, not 0**

## Usage

### Open file

```go
csv := NewCSV()
err := csv.Open("./testdata/test.csv")
if err != nil {
    panic(err)
}
defer csv.Close()
```

### Reset

You can use it reset csv file, and re-read the file. 

```go
csv.Reset()
```

after reset, file reader config will remain the same.

### Reads all rows

```go
for {
    row, isEOF, err := csv.Row()
    if isEOF {
        break
    }
    if err != nil {
        log.Fatal(err)
    }
    // Read first column data with current row, and add "A" prefix return value
    column := row.Column(1).
        TrimSpace().
        Do(func(s string) string {
            return s + "A"
        })    

    column = row.Column(1).
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
        fmt.Println(i)
    }
}
```

### Reads a row

```go
row, isEOF, err := csv.Row()
```

### Change row data

If you want to fix column value in current row, you can do it：

```go
row.Map(func (s string) string {
    return "PREFIX_" + s
}, 1)
```

The above code change first column value, will return "PREFIX_" and original column value concatenated string.

If you want change all columns value, don't pass `columnIndex` parameter value. Then all columns value will add "PREFIX_" prefix string.

### Change a column value

```go
column := row.Column(1).TrimSpace()
```

Will remove the spaces on both sides

or you can use Do() method perform custom processing, example:

```go
column := row.Column(1).Do(func(s string) string {
    if s == "a" {
        return "Number One"
    } else if s == "" {
        return "SOS"
    }
    return s
})
```

### Reads a column in the current row

```go
// Read first column in current row
column := row.Column(1)
column.TrimSpace() // Clear spaces
column.Do(func(s string) string {
	// process value
	return s
}) // Use Do method process value
column.String() // get string value

// if you want to get correct value, you will check err and continue
v, err := column.ToInt() // get int value
v, err := column.ToInt("100") // get int value, and return 100 if value is empty
v, err := column.ToFloat64() // get float value
v, err := column.ToFloat64("100.00") // get float value, and return 100.00 if value is empty
v, err := column.ToBool() // get boolean value
v, err := column.ToBool("false") // get boolean value, and return false if value is empty
v, err := column.ToTime("2006-01-02", time.Local) // get time value
v, err := column.ToTime("2006-01-02", time.Local, "2022-01-01") // get time value, and return 2022-01-01 if value is empty
```

**Valid column value conversion methods**

- String()
- ToBytes()
- ToInt()
- ToInt8()
- ToInt16()
- ToInt32()
- ToInt64()
- ToFloat32()
- ToFloat64()
- ToBool()
- ToTime()

### Save file

SaveAs() use to save a file.

This method help you create saved directory if not exists, and check you file extension save to csv/tsv/psv format. if save have any error will return it.

```go
records := make([][]string, 0)
for {
    row, isEOF, err := csvInstance.Row()
    if isEOF {
        break
    }
    if err != nil {
        log.Fatal(err)
    }
    // Do what you want to do 
    row.Map(func(s string) string {
        if row.Number != 1 {
            // Ignore header
            s = `1, "change"` + s
        }
        return s
    })
    records = append(records, row.Columns)
}
err := csvInstance.SaveAs("./a.csv", records)
```