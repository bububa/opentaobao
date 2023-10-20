package util

import (
	"sync"
)

// CrowdBindTopResultDto 结构体
type CrowdBindTopResultDto struct {
	// 新增成功的单元
	AdgroupId int64 `json:"adgroup_id,omitempty" xml:"adgroup_id,omitempty"`
}

var poolCrowdBindTopResultDto = sync.Pool{
	New: func() any {
		return new(CrowdBindTopResultDto)
	},
}

// GetCrowdBindTopResultDto() 从对象池中获取CrowdBindTopResultDto
func GetCrowdBindTopResultDto() *CrowdBindTopResultDto {
	return poolCrowdBindTopResultDto.Get().(*CrowdBindTopResultDto)
}

// ReleaseCrowdBindTopResultDto 释放CrowdBindTopResultDto
func ReleaseCrowdBindTopResultDto(v *CrowdBindTopResultDto) {
	v.AdgroupId = 0
	poolCrowdBindTopResultDto.Put(v)
}
