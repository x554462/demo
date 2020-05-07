package test

import (
	"context"
	"fmt"
	"github.com/x554462/demo/app"
	"github.com/x554462/demo/app/dao"
	"github.com/x554462/demo/app/model"
	gdao "github.com/x554462/go-dao"
	"testing"
)

func TestModel(t *testing.T) {

	app.Setup()

	districtD := dao.NewDistrictDao(gdao.NewSession(context.Background()))

	d := districtD.Select(false, 820000).(*model.District)
	fmt.Printf("%v\n", d)
	//fmt.Printf("%#v\n", d)
	//fmt.Println(d.Name.IsZero())
}
