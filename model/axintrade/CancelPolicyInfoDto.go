package axintrade

import (
	"sync"
)

// CancelPolicyInfoDto 结构体
type CancelPolicyInfoDto struct {
	// 提前小时
	Hour int64 `json:"hour,omitempty" xml:"hour,omitempty"`
	// 规则对应的值，可能是百分比、数值等
	Value int64 `json:"value,omitempty" xml:"value,omitempty"`
}

var poolCancelPolicyInfoDto = sync.Pool{
	New: func() any {
		return new(CancelPolicyInfoDto)
	},
}

// GetCancelPolicyInfoDto() 从对象池中获取CancelPolicyInfoDto
func GetCancelPolicyInfoDto() *CancelPolicyInfoDto {
	return poolCancelPolicyInfoDto.Get().(*CancelPolicyInfoDto)
}

// ReleaseCancelPolicyInfoDto 释放CancelPolicyInfoDto
func ReleaseCancelPolicyInfoDto(v *CancelPolicyInfoDto) {
	v.Hour = 0
	v.Value = 0
	poolCancelPolicyInfoDto.Put(v)
}
