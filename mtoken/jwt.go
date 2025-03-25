package mtoken

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
)

const (
	jwtAudience    = "aud"
	jwtExpire      = "exp"
	jwtID          = "jti"
	jwtIssueAt     = "iat"
	jwtIssuer      = "iss"
	jwtNotBefore   = "nbf"
	jwtSubject     = "sub"
	noDetailReason = "no detail reason"
)

// JwtMapToken jwt map token
type JwtMapToken struct {
	key []byte
}

// NewJwtMapToken 新建 JwtMapToken
func NewJwtMapToken(key []byte) *JwtMapToken {
	return &JwtMapToken{
		key: key,
	}
}

// GenerateToken 根据内容生成 jwt token
func (jmt *JwtMapToken) GenerateToken(payload jwt.MapClaims) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString(jmt.key)
	if err != nil {
		return "", fmt.Errorf("signed jwt token error:%w", err)
	}

	return token, nil
}

// ParseToken 解析 token
func (jmt *JwtMapToken) ParseToken(token string) (map[string]interface{}, error) {
	var mapClaims = make(map[string]interface{})

	jwtToken, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return jmt.key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("parse jwt claims error:%w", err)
	}
	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !jwtToken.Valid || !ok {
		return nil, errors.New("jwt claims valid error")
	}

	for k, v := range claims {
		switch k {
		case jwtAudience, jwtExpire, jwtID, jwtIssueAt, jwtIssuer, jwtNotBefore, jwtSubject:
		default:
			mapClaims[k] = v
		}
	}

	return mapClaims, nil
}
