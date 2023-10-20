package btrip

import (
	"sync"
)

// OpenCostCenterSetEntityRq 结构体
type OpenCostCenterSetEntityRq struct {
	// 人员信息列表
	EntityList []OpenOrgEntityDo `json:"entity_list,omitempty" xml:"entity_list>open_org_entity_do,omitempty"`
	// 第三方成本中心id
	ThirdpartId string `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
	// 企业id
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 商旅开放平台传2
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
}

var poolOpenCostCenterSetEntityRq = sync.Pool{
	New: func() any {
		return new(OpenCostCenterSetEntityRq)
	},
}

// GetOpenCostCenterSetEntityRq() 从对象池中获取OpenCostCenterSetEntityRq
func GetOpenCostCenterSetEntityRq() *OpenCostCenterSetEntityRq {
	return poolOpenCostCenterSetEntityRq.Get().(*OpenCostCenterSetEntityRq)
}

// ReleaseOpenCostCenterSetEntityRq 释放OpenCostCenterSetEntityRq
func ReleaseOpenCostCenterSetEntityRq(v *OpenCostCenterSetEntityRq) {
	v.EntityList = v.EntityList[:0]
	v.ThirdpartId = ""
	v.CorpId = ""
	v.Version = 0
	poolOpenCostCenterSetEntityRq.Put(v)
}
