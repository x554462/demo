package test

import (
	"context"
	"fmt"
	"github.com/x554462/demo/app/model"
	"github.com/x554462/sorm"
	"github.com/x554462/sorm/db"
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

	orm := sorm.NewSession(context.Background())

	d := orm.Get(&model.District{}).Select(false, 820000).(*model.District)
	fmt.Printf("%v\n", d)
	fmt.Printf("%#v\n", d)
}
