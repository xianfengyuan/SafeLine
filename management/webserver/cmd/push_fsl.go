package cmd

import (
	"xianfengyuan.cn/patronus/safeline-2/management/webserver/pkg/database"
	"xianfengyuan.cn/patronus/safeline-2/management/webserver/pkg/fvm"
)

func PushFSL() error {
	return fvm.PushFSL(database.GetDB().DB)
}
