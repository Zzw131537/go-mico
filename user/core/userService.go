package core

import (
	"context"
	"errors"
	"fmt"

	//"google.golang.org/genproto/googleapis/ads/googleads/v2/services"

	"user/model"
	"user/services"
)

// 进行序列化
func BuildUser(item model.User) *services.UserModel {
	userModel := &services.UserModel{
		ID:        uint32(item.ID),
		UserName:  item.UserName,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
	return userModel
}

// 用户登录
func (*UserService) UserLogin(ctx context.Context, req *services.UserRequest, resp *services.UserDetailResponse) error {
	var user model.User
	resp.Code = 200
	if err := model.DB.Where("user_name = ?", req.UserName).First(&user).Error; err != nil {
		resp.Code = 500
		return nil
	}

	if user.CheckPassWord(req.PassWord) == false {
		resp.Code = 400
		return nil
	}
	resp.UserDetail = BuildUser(user)
	return nil
}

// 用户注册
func (*UserService) UserRegister(ctx context.Context, req *services.UserRequest, resp *services.UserDetailResponse) error {

	fmt.Println("进入用户 注册函数")
	// 判断两次输入的密码是否一致
	if req.PassWord != req.PassWordConfirm {
		err := errors.New("两次输入的密码不一致")
		return err
	}
	fmt.Println(req.UserName, req.PassWord, req.PassWordConfirm)

	// 用户名唯一，判断数据库中有没有该用户
	var count int64 = 0
	if err := model.DB.Model(&model.User{}).Where("user_name=?", req.UserName).Count(&count).Error; err != nil {
		//fmt.Println("数据库查询1错误", err.Error())
		return err
	}
	if count > 0 {
		err := errors.New("该用户数据库已经存在!")
		return err
	}

	// 创建用户
	user := model.User{
		UserName: req.UserName,
	}

	// 加密密码
	if err := user.SetPassWord(req.PassWord); err != nil {
		//	fmt.Println("设置密码错误", err.Error())
		return err
	}

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		//	fmt.Println("创建用户错误!", err.Error())
		return err
	}
	// fmt.Println("注册成功")
	resp.UserDetail = BuildUser(user)
	return nil
}
