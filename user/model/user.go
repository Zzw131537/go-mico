package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	PassWordDigest string
}

const (
	PassWordCost = 12 // 密码加密难度
)

func (user *User) SetPassWord(password string) error { // 加密 密码
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PassWordDigest = string(bytes)
	return nil
}

// 检验密码
func (user *User) CheckPassWord(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PassWordDigest), []byte(password))
	return err == nil
}
