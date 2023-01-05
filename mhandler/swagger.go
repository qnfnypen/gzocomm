package mhandler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/jinzhu/copier"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/qnfnypen/gzocomm/mfile"
	"github.com/swaggo/swag"
)

var doc string

// SwaggerInfo holds exported Swagger Info so clients can modify it
type SwaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

var swaggerInfo = SwaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := swaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

// SwaggerHandler swagger 处理器
func SwaggerHandler(fp string, info *SwaggerInfo) (http.HandlerFunc, error) {
	fp = mfile.InferPathDir(fp)

	data, err := os.ReadFile(fp)
	if err != nil {
		return nil, fmt.Errorf("read swagger json file error:%w", err)
	}
	err = copier.Copy(&swaggerInfo, info)
	if err != nil {
		return nil, fmt.Errorf("copy swagger info error:%w", err)
	}
	doc = string(data)
	swag.Register(swag.Name, &s{})

	return httpSwagger.Handler(httpSwagger.URL("/swagger/doc.json")), nil
}
