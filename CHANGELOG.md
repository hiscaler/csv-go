CSV for golang Change Log
=================================

## 1.0.3 under development
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