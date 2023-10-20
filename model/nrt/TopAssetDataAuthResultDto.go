package nrt

import (
	"sync"
)

// TopAssetDataAuthResultDto 结构体
type TopAssetDataAuthResultDto struct {
	// 添加失败的手机号及原因
	Failed []AuthFailedMsg `json:"failed,omitempty" xml:"failed>auth_failed_msg,omitempty"`
}

var poolTopAssetDataAuthResultDto = sync.Pool{
	New: func() any {
		return new(TopAssetDataAuthResultDto)
	},
}

// GetTopAssetDataAuthResultDto() 从对象池中获取TopAssetDataAuthResultDto
func GetTopAssetDataAuthResultDto() *TopAssetDataAuthResultDto {
	return poolTopAssetDataAuthResultDto.Get().(*TopAssetDataAuthResultDto)
}

// ReleaseTopAssetDataAuthResultDto 释放TopAssetDataAuthResultDto
func ReleaseTopAssetDataAuthResultDto(v *TopAssetDataAuthResultDto) {
	v.Failed = v.Failed[:0]
	poolTopAssetDataAuthResultDto.Put(v)
}
