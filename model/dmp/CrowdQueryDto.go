package dmp

import (
	"sync"
)

// CrowdQueryDto 结构体
type CrowdQueryDto struct {
	// 按已有人群id查询过滤
	CrowdIdList []int64 `json:"crowd_id_list,omitempty" xml:"crowd_id_list>int64,omitempty"`
	// 按人群名称模糊匹配
	CrowdName string `json:"crowd_name,omitempty" xml:"crowd_name,omitempty"`
}

var poolCrowdQueryDto = sync.Pool{
	New: func() any {
		return new(CrowdQueryDto)
	},
}

// GetCrowdQueryDto() 从对象池中获取CrowdQueryDto
func GetCrowdQueryDto() *CrowdQueryDto {
	return poolCrowdQueryDto.Get().(*CrowdQueryDto)
}

// ReleaseCrowdQueryDto 释放CrowdQueryDto
func ReleaseCrowdQueryDto(v *CrowdQueryDto) {
	v.CrowdIdList = v.CrowdIdList[:0]
	v.CrowdName = ""
	poolCrowdQueryDto.Put(v)
}
