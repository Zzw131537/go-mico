package core

import (
	"context"
	"errors"
	"fmt"
	"user/model"
	"user/service"

	"gorm.io/gorm"
)

func BuildUser(user *model.User) *service.UserModel {
	return &service.UserModel{
		ID:        uint32(user.ID),
		UserName:  user.UserName,
		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: user.UpdatedAt.Unix(),
	}
}

func (*UserService) UserLogin(ctx context.Context, req *service.UserRequest, resp *service.UserDetailResponse) error {
	var user model.User

	resp.Code = 200
	if err := model.DB.Where("user_name = ?", req.UserName).First(&user).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			resp.Code = 400
			return nil
		}
		resp.Code = 500
		return nil
	}
	if user.CheckPassword(req.PassWord) == false {
		resp.Code = 400
		return nil
	}
	resp.UserDetail = BuildUser(&user)
	return nil

}

func (*UserService) UserRegister(ctx context.Context, req *service.UserRequest, resp *service.UserDetailResponse) error {
	if req.PassWord != req.PassWordConfirm {
		err := errors.New("两次密码输入不一致")
		return err
	}
	fmt.Println("用户名: " + req.UserName + "密码: " + req.PassWord)
	var count int64
	count = 0
	if err := model.DB.Model(&model.User{}).Where("user_name=?", req.UserName).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户名已经存在")
	}
	user := &model.User{
		UserName: req.UserName,
	}
	if err := user.SetPassword(req.PassWord); err != nil {
		return nil
	}
	if err := model.DB.Create(&user).Error; err != nil {
		return err
	}
	resp.UserDetail = BuildUser(user)
	return nil

}
