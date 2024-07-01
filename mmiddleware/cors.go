package mmiddleware

import (
	"net/http"
)

// CorsMiddleware 跨域请求处理中间件
type CorsMiddleware struct{}

// NewCorsMiddleware 新建跨域请求处理中间件
func NewCorsMiddleware() *CorsMiddleware {
	return &CorsMiddleware{}
}

// Handle 跨域请求处理
func (m *CorsMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 设置响应头
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// 放行所有 OPTIONS 方法
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// 处理请求
		next(w, r)
	}
}
