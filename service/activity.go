package service

import (
	"net/http"

	"github.com/jinzhu/copier"

	"customer/db/mysql"
	"customer/model"
	"customer/request"
	"customer/response"

	"gorm.io/gorm"
)

func CreateActivity(req request.CreateActivityReq) (response.Response, error) {
	var activity model.Activity
	copier.Copy(&activity, &req)
	err := mysql.MysqlDB.Create(&activity).Error
	if err != nil {
		return response.Response{
			Status: http.StatusInternalServerError,
			Msg:    "数据库添加数据出错！",
			Error:  err.Error(),
		}, err
	}
	return response.Response{
		Status: http.StatusOK,
		Msg:    "活动新增成功！",
	}, err
}

// 更新活动
func UpdateActivity(req request.UpdateActivityReq) (response.Response, error) {
	var activity model.Activity
	result := mysql.MysqlDB.First(&model.Activity{}, req.ID)
	if result.Error == gorm.ErrRecordNotFound { // ErrRecordNotFound错说明数据库没有记录
		return response.Response{
			Status: http.StatusInternalServerError,
			Msg:    "数据库无此数据！",
			Error:  result.Error.Error(),
		}, result.Error
	} else if result.Error != nil {
		return response.Response{
			Status: http.StatusInternalServerError,
			Msg:    "数据库内部出错！",
			Error:  result.Error.Error(),
		}, result.Error
	}

	copier.Copy(&activity, &req)
	err := mysql.MysqlDB.Save(&activity).Error
	if err != nil {
		return response.Response{
			Status: http.StatusInternalServerError,
			Msg:    "数据库更新数据出错！",
			Error:  err.Error(),
		}, err
	}
	return response.Response{
		Status: http.StatusOK,
		Msg:    "活动更新成功！",
	}, err
}

// 根据ID获取活动
func GetActivityByID(req request.GetActivityReq) (response.Response, error) {
	var activity model.Activity
	result := mysql.MysqlDB.First(&activity, req.ID)
	if result.Error == gorm.ErrRecordNotFound { // ErrRecordNotFound错说明数据库没有记录
		return response.Response{
			Status: http.StatusInternalServerError,
			Msg:    "数据库无此数据！",
			Error:  result.Error.Error(),
		}, result.Error
	} else if result.Error != nil {
		return response.Response{
			Status: http.StatusInternalServerError,
			Msg:    "数据库内部出错！",
			Error:  result.Error.Error(),
		}, result.Error
	}

	var respActivity model.RespActivity
	copier.Copy(&respActivity, &activity)
	return response.Response{
		Status: http.StatusOK,
		Msg:    "获取活动成功！",
		Data:   respActivity,
	}, nil
}
