package api

import (
	"net/http"

	"customer/request"
	"customer/service"

	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func CreateActivity(c *gin.Context) {
	var createActivityReq request.CreateActivityReq
	err := c.ShouldBind(&createActivityReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"msg":    "json数据解析失败！",
			"error":  err.Error(),
		})
		logging.Info(err)
		return
	}
	res, err := service.CreateActivity(createActivityReq)
	c.JSON(res.Status, res)
}

func UpdateActivity(c *gin.Context) {
	var updateActivity request.UpdateActivityReq
	err := c.ShouldBind(&updateActivity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"msg":    "json数据解析失败！",
			"error":  err.Error(),
		})
		logging.Info(err)
		return
	}
	res, err := service.UpdateActivity(updateActivity)
	c.JSON(res.Status, res)
}

//更新ID获取活动
func GetActivityByID(c *gin.Context) {
	var getActivity request.GetActivityReq
	err := c.ShouldBind(&getActivity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"msg":    "json数据解析失败！",
			"error":  err.Error(),
		})
		logging.Info(err)
		return
	}
	res, err := service.GetActivityByID(getActivity)
	c.JSON(res.Status, res)
}
