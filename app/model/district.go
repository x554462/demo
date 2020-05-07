package model

import (
	"database/sql"
	"github.com/x554462/go-dao/excode"
	"github.com/x554462/go-exception"
)

var DistrictNotFoundError = exception.New("地区未找到", excode.ModelNotFoundError)

type District struct {
	Code         int             `db:"code"`
	Father       sql.NullInt64   `db:"father"`
	Name         sql.NullString  `db:"name"`
	ShortName    sql.NullString  `db:"short_name"`
	Type         sql.NullInt32   `db:"type"`
	EnglishName  sql.NullString  `db:"english_name"`
	Abbreviation sql.NullString  `db:"abbreviation"`
	Longitude    sql.NullFloat64 `db:"longitude"`
	Latitude     sql.NullFloat64 `db:"latitude"`
}

func (this *District) GetIndexValues() []interface{} {
	return []interface{}{this.Code}
}

func (this *District) InitModelInfo() (tableName string, indexFields []string, notFoundErr exception.ErrorWrapper) {
	return "district", []string{"code"}, DistrictNotFoundError
}
