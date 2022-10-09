package mfile

import (
	"os"
	"path/filepath"
	"strings"
)

// CheckFileExist 判断给定的文件目录是否存在
func CheckFileExist(fp string) bool {
	_, err := os.Stat(fp)

	return err == nil || os.IsExist(err)
}

// InferPathDir 获取包含此目录的最近的文件夹，错误则返回空
func InferPathDir(sp string) string {
	sp = "/" + strings.TrimPrefix(sp, "/")
	// 获取当前目录
	pwd, err := os.Getwd()
	if err != nil {
		return ""
	}

	var infer func(d string) string
	infer = func(d string) string {
		if d == "/" {
			return ""
		}

		if CheckFileExist(d + sp) {
			return d
		}

		return infer(filepath.Dir(d))
	}

	return infer(pwd) + sp
}
