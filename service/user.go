package service

import (
	"net/http"

	"customer/db/mysql"
	"customer/model"
	"customer/request"
	"customer/response"

	"github.com/jinzhu/copier"
)

func UserJoin(newUser request.UserReq) (response.Response, error) {
	var user model.User
	var activity model.Activity
	err := mysql.MysqlDB.First(&activity, newUser.Activity).Error
	if err != nil {
		return response.Response{
			Status: http.StatusInternalServerError,
			Msg:    "数据库出错！",
			Error:  err.Error(),
		}, err
	}

	if activity.Current == activity.Total {
		return response.Response{
			Status: http.StatusBadGateway,
			Msg:    "报名人数已满！",
			Error:  "报名人数已满！",
		}, err
	}

	tx := mysql.MysqlDB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 更新活动已报名人数
	activity.Current = activity.Current + 1
	err = tx.Save(&activity).Error
	if err != nil {
		tx.Rollback()
		return response.Response{
			Status: http.StatusInternalServerError,
			Msg:    "数据库更新数据出错！",
			Error:  err.Error(),
		}, err
	}

	copier.Copy(&user, &newUser)
	err = tx.Create(&user).Error
	if err != nil {
		tx.Rollback()

		// 更新活动已报名人数
		activity.Current = activity.Current - 1
		err = tx.Save(&activity).Error

		return response.Response{
			Status: http.StatusInternalServerError,
			Msg:    "数据库添加数据出错！",
			Error:  err.Error(),
		}, err

	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return response.Response{
			Status: http.StatusInternalServerError,
			Msg:    "数据库事务提交失败！",
			Error:  err.Error(),
		}, err
	}

	return response.Response{
		Status: http.StatusOK,
		Msg:    "用户报名成功！",
	}, err
}
