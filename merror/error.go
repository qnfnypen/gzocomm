package merror

import (
	"fmt"
	"path"
	"runtime"
	"time"
)

// GzErr 自定义错误
type GzErr struct {
	LogErr  string
	RespErr string
}

// NewErr 自定义错误
type NewErr struct {
	FuncName  string
	File      string
	TimeStamp string
	Err       string
}

// Error 实现 error 接口
func (gz GzErr) Error() string {
	return gz.RespErr
}

// Error 实现 error 接口
func (ne NewErr) Error() string {
	return fmt.Sprintf("func:%s file:%s timestamp:%v error:%v", ne.FuncName, ne.File, time.Now().Format("2006-01-02 15:04:05"), ne.Err)
}

// NewError 根据 error 新建 error
func NewError(err error) error {
	var ne NewErr
	// 判断是否已经为新错误
	if ne, ok := err.(NewErr); ok {
		return ne
	}
	// 获取函数名和行数
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return err
	}
	funcName := runtime.FuncForPC(pc).Name()
	ne.FuncName = path.Base(funcName)
	ne.File = fmt.Sprintf("%s:%d", file, line)
	ne.TimeStamp = time.Now().Format("2006-01-02 15:04:05")
	ne.Err = err.Error()

	return ne
}
