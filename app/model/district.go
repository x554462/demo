package model

import (
	"database/sql"
	"github.com/x554462/go-exception"
	"github.com/x554462/sorm"
)

var DistrictNotFoundError = exception.New("地区未找到", sorm.ModelNotFoundError)

type District struct {
	sorm.BaseModel
	Code         sql.NullInt64   `db:"code,pk"`
	Father       sql.NullInt64   `db:"father"`
	Name         sql.NullString  `db:"name"`
	ShortName    sql.NullString  `db:"short_name"`
	Type         sql.NullInt32   `db:"type"`
	EnglishName  sql.NullString  `db:"english_name"`
	Abbreviation sql.NullString  `db:"abbreviation"`
	Longitude    sql.NullFloat64 `db:"longitude"`
	Latitude     sql.NullFloat64 `db:"latitude"`
}

func (this *District) IndexValues() []interface{} {
	return []interface{}{this.Code}
}

func (this *District) GetNotFoundError() exception.ErrorWrapper {
	return DistrictNotFoundError
}
