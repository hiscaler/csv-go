CSV For Golang
==============

csv-go is csv file helper function, support csv, tsv format

## Install

```go
go get github.com/hiscaler/csv-go
```

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
```

### Read a row
```go
row, err := csv.Read()
if err == io.EOF {
    break
}
if err != nil {
    log.Fatal(err)
}
```

### Read row column
```go
// Read first column in current row
value := row.Read(0)
value.TrimSpace() // Clear spaces
value.Do(func(s string) string {
	// process value
	return s
}) // Use Do method process value
value.String() // get value with string

// if you want to get correct value, you will check err and continue
v, err := value.ToInt() // get int value
v, err := value.ToInt("100") // get int value, and return 100 if value is empty
v, err := value.ToFloat64() // get float value
v, err := value.ToFloat64("100.00") // get float value, and return 100.00 if value is empty
v, err := value.ToBool() // get boolean value
v, err := value.ToBool("false") // get boolean value, and return false if value is empty
v, err := value.ToTime() // get time value
v, err := value.ToTime("2022-01-01") // get time value, and return 2022-01-01 if value is empty
```

## Notice

1. **row start index is 1**
2. **column start index is 0**