package alilabs

import (
	"sync"
)

// TopRefreshReqDto 结构体
type TopRefreshReqDto struct {
	// clientId
	ClientId string `json:"client_id,omitempty" xml:"client_id,omitempty"`
	// 只支持“basic”
	Scope string `json:"scope,omitempty" xml:"scope,omitempty"`
	// 只支持“refresh_token”
	GrantType string `json:"grant_type,omitempty" xml:"grant_type,omitempty"`
	// refreshToken
	RefreshToken string `json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
	// 设备uuid
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}

var poolTopRefreshReqDto = sync.Pool{
	New: func() any {
		return new(TopRefreshReqDto)
	},
}

// GetTopRefreshReqDto() 从对象池中获取TopRefreshReqDto
func GetTopRefreshReqDto() *TopRefreshReqDto {
	return poolTopRefreshReqDto.Get().(*TopRefreshReqDto)
}

// ReleaseTopRefreshReqDto 释放TopRefreshReqDto
func ReleaseTopRefreshReqDto(v *TopRefreshReqDto) {
	v.ClientId = ""
	v.Scope = ""
	v.GrantType = ""
	v.RefreshToken = ""
	v.Uuid = ""
	poolTopRefreshReqDto.Put(v)
}
