package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/x554462/demo/app/dao"
	"github.com/x554462/demo/middleware/mango"
)

func GetTest(c *gin.Context) {
	mg := mango.Default(c)

	districtD := dao.NewDistrictDao(mg.GetDaoSession())

	d := districtD.Select(false, 1)
	fmt.Println(d)
}
