package jym

import (
	"sync"
)

// RealNameVerifyTopDto 结构体
type RealNameVerifyTopDto struct {
	// 实名校验结果信息
	VerifyMsg string `json:"verify_msg,omitempty" xml:"verify_msg,omitempty"`
	// 实名校验结果信息
	IdentityCode string `json:"identity_code,omitempty" xml:"identity_code,omitempty"`
	// 实名校验结果信息
	VerifyCode int64 `json:"verify_code,omitempty" xml:"verify_code,omitempty"`
}

var poolRealNameVerifyTopDto = sync.Pool{
	New: func() any {
		return new(RealNameVerifyTopDto)
	},
}

// GetRealNameVerifyTopDto() 从对象池中获取RealNameVerifyTopDto
func GetRealNameVerifyTopDto() *RealNameVerifyTopDto {
	return poolRealNameVerifyTopDto.Get().(*RealNameVerifyTopDto)
}

// ReleaseRealNameVerifyTopDto 释放RealNameVerifyTopDto
func ReleaseRealNameVerifyTopDto(v *RealNameVerifyTopDto) {
	v.VerifyMsg = ""
	v.IdentityCode = ""
	v.VerifyCode = 0
	poolRealNameVerifyTopDto.Put(v)
}
