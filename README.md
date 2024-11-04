# go-sql-raw
![Tests](https://github.com/mysiar-org/go_sql_raw/actions/workflows/tests.yml/badge.svg)
[![Gosec](https://github.com/mysiar-org/go_sql_raw/actions/workflows/gosec.yml/badge.svg)](https://github.com/mysiar-org/go_sql_raw/actions/workflows/gosec.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/mysiar-org/go_sql_raw.svg)](https://pkg.go.dev/github.com/mysiar-org/go-sql-raw)
![GitHub go.mod Go version (branch)](https://img.shields.io/github/go-mod/go-version/mysiar-org/go_sql_raw/main)


The main reason that this module has been created was generic querying database without knowing table structure. 

## How to use it

Install
```bash
go get github.com/mysiar-org/go-sql-raw
```

Example use
```go
rows, err: = Db.Query("SELECT * FROM album")

defer rows.Close()
var data []go_sql_raw.RawSqlType
for rows.Next() {
	rec := go_sql_raw.Rows2Map(rows)
	data = append(data, rec)
}
```

check also [go_sql_raw_test.go](tests/go_sql_raw_test.go)

## TODO
* update type converting from DB to go in function `convertType`

## Credits
* module inspired by https://gist.github.com/SchumacherFM/69a167bec7dea644a20e
