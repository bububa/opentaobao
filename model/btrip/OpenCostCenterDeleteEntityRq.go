package btrip

import (
	"sync"
)

// OpenCostCenterDeleteEntityRq 结构体
type OpenCostCenterDeleteEntityRq struct {
	// 删除的成员信息列表,del_all为true时可不填
	EntityList []OpenOrgEntityDo `json:"entity_list,omitempty" xml:"entity_list>open_org_entity_do,omitempty"`
	// 第三方成本中心id
	ThirdpartId string `json:"thirdpart_id,omitempty" xml:"thirdpart_id,omitempty"`
	// 企业id
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 商旅开放平台传2
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
	// 是否全部删除
	DelAll bool `json:"del_all,omitempty" xml:"del_all,omitempty"`
}

var poolOpenCostCenterDeleteEntityRq = sync.Pool{
	New: func() any {
		return new(OpenCostCenterDeleteEntityRq)
	},
}

// GetOpenCostCenterDeleteEntityRq() 从对象池中获取OpenCostCenterDeleteEntityRq
func GetOpenCostCenterDeleteEntityRq() *OpenCostCenterDeleteEntityRq {
	return poolOpenCostCenterDeleteEntityRq.Get().(*OpenCostCenterDeleteEntityRq)
}

// ReleaseOpenCostCenterDeleteEntityRq 释放OpenCostCenterDeleteEntityRq
func ReleaseOpenCostCenterDeleteEntityRq(v *OpenCostCenterDeleteEntityRq) {
	v.EntityList = v.EntityList[:0]
	v.ThirdpartId = ""
	v.CorpId = ""
	v.Version = 0
	v.DelAll = false
	poolOpenCostCenterDeleteEntityRq.Put(v)
}
