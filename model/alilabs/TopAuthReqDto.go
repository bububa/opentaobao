package alilabs

import (
	"sync"
)

// TopAuthReqDto 结构体
type TopAuthReqDto struct {
	// 二维码授权 只支持qrcode
	ResponseType string `json:"response_type,omitempty" xml:"response_type,omitempty"`
	// 天猫精灵分配的clientId
	ClientId string `json:"client_id,omitempty" xml:"client_id,omitempty"`
	// OAUTH2 scope 只支持basic
	Scope string `json:"scope,omitempty" xml:"scope,omitempty"`
	// OAUTH2 state 随意填写
	State string `json:"state,omitempty" xml:"state,omitempty"`
}

var poolTopAuthReqDto = sync.Pool{
	New: func() any {
		return new(TopAuthReqDto)
	},
}

// GetTopAuthReqDto() 从对象池中获取TopAuthReqDto
func GetTopAuthReqDto() *TopAuthReqDto {
	return poolTopAuthReqDto.Get().(*TopAuthReqDto)
}

// ReleaseTopAuthReqDto 释放TopAuthReqDto
func ReleaseTopAuthReqDto(v *TopAuthReqDto) {
	v.ResponseType = ""
	v.ClientId = ""
	v.Scope = ""
	v.State = ""
	poolTopAuthReqDto.Put(v)
}
