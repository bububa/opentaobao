package btrip

import (
	"sync"
)

// OpenCostCenterDeleteRq 结构体
type OpenCostCenterDeleteRq struct {
	// 第三方成本中心id
	ThirdpartId string `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
	// 企业id
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 商旅开放平台传2
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
}

var poolOpenCostCenterDeleteRq = sync.Pool{
	New: func() any {
		return new(OpenCostCenterDeleteRq)
	},
}

// GetOpenCostCenterDeleteRq() 从对象池中获取OpenCostCenterDeleteRq
func GetOpenCostCenterDeleteRq() *OpenCostCenterDeleteRq {
	return poolOpenCostCenterDeleteRq.Get().(*OpenCostCenterDeleteRq)
}

// ReleaseOpenCostCenterDeleteRq 释放OpenCostCenterDeleteRq
func ReleaseOpenCostCenterDeleteRq(v *OpenCostCenterDeleteRq) {
	v.ThirdpartId = ""
	v.CorpId = ""
	v.Version = 0
	poolOpenCostCenterDeleteRq.Put(v)
}
