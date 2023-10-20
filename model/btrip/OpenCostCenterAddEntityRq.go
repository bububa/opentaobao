package btrip

import (
	"sync"
)

// OpenCostCenterAddEntityRq 结构体
type OpenCostCenterAddEntityRq struct {
	// 人员信息列表
	EntityList []OpenOrgEntityDo `json:"entity_list,omitempty" xml:"entity_list>open_org_entity_do,omitempty"`
	// 第三方成本中心id
	ThirdpartId string `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
	// 企业id
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 商旅开放平台传2
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
}

var poolOpenCostCenterAddEntityRq = sync.Pool{
	New: func() any {
		return new(OpenCostCenterAddEntityRq)
	},
}

// GetOpenCostCenterAddEntityRq() 从对象池中获取OpenCostCenterAddEntityRq
func GetOpenCostCenterAddEntityRq() *OpenCostCenterAddEntityRq {
	return poolOpenCostCenterAddEntityRq.Get().(*OpenCostCenterAddEntityRq)
}

// ReleaseOpenCostCenterAddEntityRq 释放OpenCostCenterAddEntityRq
func ReleaseOpenCostCenterAddEntityRq(v *OpenCostCenterAddEntityRq) {
	v.EntityList = v.EntityList[:0]
	v.ThirdpartId = ""
	v.CorpId = ""
	v.Version = 0
	poolOpenCostCenterAddEntityRq.Put(v)
}
