package scbp

import (
	"sync"
)

// CrowdInfo 结构体
type CrowdInfo struct {
	// 人群预估数量
	EstimateCountList []int64 `json:"estimate_count_list,omitempty" xml:"estimate_count_list>int64,omitempty"`
	// 人群中文名称
	CrowdName string `json:"crowd_name,omitempty" xml:"crowd_name,omitempty"`
	// 人群id
	CrowdId string `json:"crowd_id,omitempty" xml:"crowd_id,omitempty"`
}

var poolCrowdInfo = sync.Pool{
	New: func() any {
		return new(CrowdInfo)
	},
}

// GetCrowdInfo() 从对象池中获取CrowdInfo
func GetCrowdInfo() *CrowdInfo {
	return poolCrowdInfo.Get().(*CrowdInfo)
}

// ReleaseCrowdInfo 释放CrowdInfo
func ReleaseCrowdInfo(v *CrowdInfo) {
	v.EstimateCountList = v.EstimateCountList[:0]
	v.CrowdName = ""
	v.CrowdId = ""
	poolCrowdInfo.Put(v)
}
