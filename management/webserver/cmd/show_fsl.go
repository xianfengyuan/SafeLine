package cmd

import (
	"xianfengyuan.cn/patronus/safeline-2/management/webserver/pkg/database"
	"xianfengyuan.cn/patronus/safeline-2/management/webserver/pkg/fvm"
)

func ShowFSL() (string, error) {
	return fvm.GenerateFullFSL(database.GetDB().DB)
}
