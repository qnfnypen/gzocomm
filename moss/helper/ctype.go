package helper

import (
	"io"
	"mime"
	"net/http"
)

// GetContentTypeAndExt 获取body的content-type和ext
func GetContentTypeAndExt(f io.Reader, ctype, ext string) (string, string, error) {
	// 根据 ext 获取 content-type
	if ext != "" && ctype == "" {
		ctype = mime.TypeByExtension(ext)
	}
	if ctype == "" {
		var buf [512]byte
		n, err := io.ReadFull(f, buf[:])
		if err != nil {
			return "", "", err
		}
		ctype = http.DetectContentType(buf[:n])
	}
	// 根据 content-type 获取 ext
	if ext == "" {
		exts, err := mime.ExtensionsByType(ctype)
		if err != nil {
			return "", "", err
		}
		if exts != nil {
			ext = exts[0]
		}
	}

	return ctype, ext, nil
}
