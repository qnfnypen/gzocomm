package muuid

import "github.com/rs/xid"

// GenerateUUID 生成唯一标识
func GenerateUUID() string {
	guid := xid.New()

	return guid.String()
}
