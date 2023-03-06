package response

import (
	"fmt"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// Body 响应体
type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Time string      `json:"time"`
	Data interface{} `json:"data,omitempty"`
}

// Response 重写Response方法
func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err != nil {
		body.Code = -1
		body.Msg = err.Error()
	} else {
		body.Msg = "OK"
		body.Data = resp
	}
	body.Time = time.Now().Format("2006-01-02 15:04:05.000")
	httpx.OkJson(w, body)
}

// ParamError 参数错误
func ParamError(w http.ResponseWriter, err error) {
	var body Body

	body.Code = -1
	body.Msg = fmt.Sprintf("param error:%v", err.Error())

	body.Time = time.Now().Format("2006-01-02 15:04:05.000")
	httpx.OkJson(w, body)
}

// UnauthorizedError 授权过期
func UnauthorizedError(w http.ResponseWriter, errMsg string) {
	var body Body

	body.Code = 401
	body.Msg = errMsg

	body.Time = time.Now().Format("2006-01-02 15:04:05.000")
	httpx.OkJson(w, body)
}

// CodeResponse 指定
func CodeResponse(w http.ResponseWriter, resp interface{}, code int, err error) {
	var body Body

	if err != nil {
		body.Code = code
		body.Msg = err.Error()
	} else {
		body.Msg = "OK"
		body.Data = resp
	}

	body.Time = time.Now().Format("2006-01-02 15:04:05.000")
	httpx.OkJson(w, body)
}
