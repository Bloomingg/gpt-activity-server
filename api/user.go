package api

import (
	"customer/request"
	"customer/service"
	"net/http"

	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func UserJoin(c *gin.Context) {
	var userJoin request.UserReq
	err := c.ShouldBind(&userJoin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"msg":    "json数据解析失败！",
			"error":  err.Error(),
		})
		logging.Info(err)
		return
	}
	res, err := service.UserJoin(userJoin)
	c.JSON(res.Status, res)
}
