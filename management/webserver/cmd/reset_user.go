package cmd

import (
	"xianfengyuan.cn/patronus/safeline-2/management/webserver/model"
	"xianfengyuan.cn/patronus/safeline-2/management/webserver/pkg/database"
)

func ResetUser(username string) {
	db := database.GetDB()
	var user model.User
	db.Where(&model.User{Username: username}).First(&user)
	user.LastLoginTime = 0
	db.Save(&user)
}
