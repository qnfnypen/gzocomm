package merror

// GzErr 自定义错误
type GzErr struct {
	LogErr  string
	RespErr string
}

// Error 实现 error 接口
func (gz *GzErr) Error() string {
	return gz.RespErr
}
