package dao

import (
	"github.com/x554462/demo/app/model"
	"github.com/x554462/go-dao"
)

type DistrictDao struct {
	dao.Dao
}

func NewDistrictDao(ds *dao.Session) *DistrictDao {
	return ds.BindOnce(&model.District{}, &DistrictDao{}).(*DistrictDao)
}
