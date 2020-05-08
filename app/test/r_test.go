package test

import (
	"context"
	"fmt"
	"github.com/x554462/demo/app/dao"
	"github.com/x554462/demo/app/model"
	gdao "github.com/x554462/go-dao"
	"github.com/x554462/go-dao/db"
	"testing"
)

func TestModel(t *testing.T) {

	db.Setup(db.Conf{
		Name:     "db",
		User:     "root",
		Password: "123456",
		Host:     "127.0.0.1",
		Port:     3306,
	})

	districtD := dao.NewDistrictDao(gdao.NewSession(context.Background()))

	d := districtD.Select(false, 820000).(*model.District)
	fmt.Printf("%v\n", d)
	//fmt.Printf("%#v\n", d)
	//fmt.Println(d.Name.IsZero())
}
