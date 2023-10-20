package util

import (
	"sync"
)

// CrowdTopDto 结构体
type CrowdTopDto struct {
	// 达摩盘生成的人群id
	CrowdValue string `json:"crowd_value,omitempty" xml:"crowd_value,omitempty"`
	// 传对应的人群名称
	CrowdName string `json:"crowd_name,omitempty" xml:"crowd_name,omitempty"`
	// 达摩盘人群传128，平台精选人群传4194304
	TargetType int64 `json:"target_type,omitempty" xml:"target_type,omitempty"`
}

var poolCrowdTopDto = sync.Pool{
	New: func() any {
		return new(CrowdTopDto)
	},
}

// GetCrowdTopDto() 从对象池中获取CrowdTopDto
func GetCrowdTopDto() *CrowdTopDto {
	return poolCrowdTopDto.Get().(*CrowdTopDto)
}

// ReleaseCrowdTopDto 释放CrowdTopDto
func ReleaseCrowdTopDto(v *CrowdTopDto) {
	v.CrowdValue = ""
	v.CrowdName = ""
	v.TargetType = 0
	poolCrowdTopDto.Put(v)
}
