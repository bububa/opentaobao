package alihouse

import (
	"sync"
)

// IdentityInfoDto 结构体
type IdentityInfoDto struct {
	// 经纪人ID
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 身份信息
	Identity string `json:"identity,omitempty" xml:"identity,omitempty"`
	// 版本号
	EtcVersion int64 `json:"etc_version,omitempty" xml:"etc_version,omitempty"`
	// 类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolIdentityInfoDto = sync.Pool{
	New: func() any {
		return new(IdentityInfoDto)
	},
}

// GetIdentityInfoDto() 从对象池中获取IdentityInfoDto
func GetIdentityInfoDto() *IdentityInfoDto {
	return poolIdentityInfoDto.Get().(*IdentityInfoDto)
}

// ReleaseIdentityInfoDto 释放IdentityInfoDto
func ReleaseIdentityInfoDto(v *IdentityInfoDto) {
	v.OuterId = ""
	v.Identity = ""
	v.EtcVersion = 0
	v.Type = 0
	poolIdentityInfoDto.Put(v)
}
