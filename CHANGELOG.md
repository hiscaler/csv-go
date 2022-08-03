CSV for golang Change Log
=================================

## 1.0.7 under development

- Chg: Fixed demo code error in readme file
- Enh: Optimize Map method
- Bug: Fixed column.TrimSpace and Do method

## 1.0.6

- Chg: Set reader to public, user manually sets the parameters
- Chg: Row and Column start index from 1

## 1.0.5

- Chg: IsBlack() rename to IsBlank()
- New: Add IsNull() method for column
- Chg: Fixed Reset method and support BOM file

## 1.0.4

- New: Add ToBytes() method for column
- New: Add csv.Reset() method for re-read file

## 1.0.3

- Chg: CSV.Row() return variable name isLastRow change to isEOF
- Enh: If CSV.Row() isEOF return true, then err value force set to nil

## 1.0.2

- New: Add csv.Close() method for close open file
- New: Support save to TSV format
- Enh: SaveAs method return error if write failed
- Enh: Perfect digital string processing

## 1.0.1

- Chg: Perfect doc
- Chg: csv.Open() method remove header and data rows number parameters
- New: Add IsEmpty() method for row
- New: Support file save
- Chg: Remove row.Record() method, you can get from row.Columns attribute

## 1.0.0

- Initial release.