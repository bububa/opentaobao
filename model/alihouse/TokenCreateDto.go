package alihouse

import (
	"sync"
)

// TokenCreateDto 结构体
type TokenCreateDto struct {
	// 授权码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否测试(1-是测试 0-不是测试)
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
}

var poolTokenCreateDto = sync.Pool{
	New: func() any {
		return new(TokenCreateDto)
	},
}

// GetTokenCreateDto() 从对象池中获取TokenCreateDto
func GetTokenCreateDto() *TokenCreateDto {
	return poolTokenCreateDto.Get().(*TokenCreateDto)
}

// ReleaseTokenCreateDto 释放TokenCreateDto
func ReleaseTokenCreateDto(v *TokenCreateDto) {
	v.Code = ""
	v.IsTest = 0
	poolTokenCreateDto.Put(v)
}
