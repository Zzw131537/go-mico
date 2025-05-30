package model

import (
	"fmt"
)

// 数据库迁移
func migration() {
	err := DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		&Task{},
	)
	if err != nil {
		fmt.Println("err ", err)
	}
	return
}
