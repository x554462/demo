package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/x554462/demo/app/model"
	"github.com/x554462/demo/middleware/mango"
)

func GetTest(c *gin.Context) {
	mg := mango.Default(c)

	d := mg.GetOrmSession().GetDao(&model.District{}).Select(false, 1)
	fmt.Println(d)
}
