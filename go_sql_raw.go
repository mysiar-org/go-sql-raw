package go_sql_raw

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

type RawSqlType map[string]interface{}

func Error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Rows2Map(rows *sql.Rows) RawSqlType {
	columns, err := rows.Columns()
	columnTypes, err := rows.ColumnTypes()
	Error(err)
	rc := newMapRawSqlScan(columns, columnTypes)
	err = rc.update(rows)
	Error(err)
	return rc.get()
}

type mapRawSqlScan struct {
	cp       []interface{}
	row      RawSqlType
	colCount int
	colNames []string
	colTypes []*sql.ColumnType
}

func (s *mapRawSqlScan) update(rows *sql.Rows) error {
	if err := rows.Scan(s.cp...); err != nil {
		return err
	}

	for i := 0; i < s.colCount; i++ {
		if rb, ok := s.cp[i].(*sql.RawBytes); ok {
			val := string(*rb)
			var parsed any
			switch s.colTypes[i].DatabaseTypeName() {
			case "DECIMAL":
				parsed, _ = strconv.ParseFloat(val, 64)
			case "INT":
				parsed, _ = strconv.ParseInt(val, 10, 64)
			default:
				parsed = val
			}
			s.row[s.colNames[i]] = parsed
			*rb = nil
		} else {
			return fmt.Errorf("Cannot convert index %d column %s to type *sql.RawBytes", i, s.colNames[i])
		}
	}

	return nil
}

func (s *mapRawSqlScan) get() RawSqlType {
	return s.row
}

func newMapRawSqlScan(columnNames []string, columnTypes []*sql.ColumnType) *mapRawSqlScan {
	lenCN := len(columnNames)
	s := &mapRawSqlScan{
		cp:       make([]interface{}, lenCN),
		row:      make(RawSqlType, lenCN),
		colCount: lenCN,
		colNames: columnNames,
		colTypes: columnTypes,
	}
	for i := 0; i < lenCN; i++ {
		s.cp[i] = new(sql.RawBytes)
	}
	return s
}
