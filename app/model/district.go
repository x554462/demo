package model

import (
	"github.com/x554462/go-exception"
	"github.com/x554462/sorm"
	"github.com/x554462/sorm/type"
)

var DistrictNotFoundError = exception.New("地区未找到", sorm.ModelNotFoundError)

type District struct {
	sorm.BaseModel
	Code         int          `db:"code,pk"`
	Father       _type.Int    `db:"father"`
	Name         _type.String `db:"name"`
	ShortName    _type.String `db:"short_name"`
	Type         _type.Int    `db:"type"`
	EnglishName  _type.String `db:"english_name"`
	Abbreviation _type.String `db:"abbreviation"`
	Longitude    _type.Float  `db:"longitude"`
	Latitude     _type.Float  `db:"latitude"`
}

func (this *District) IndexValues() []interface{} {
	return []interface{}{this.Code}
}

func (this *District) GetNotFoundError() exception.ErrorWrapper {
	return DistrictNotFoundError
}
