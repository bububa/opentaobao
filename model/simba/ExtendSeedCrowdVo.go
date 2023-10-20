package simba

import (
	"sync"
)

// ExtendSeedCrowdVo 结构体
type ExtendSeedCrowdVo struct {
	// 人群名称
	CrowdName string `json:"crowd_name,omitempty" xml:"crowd_name,omitempty"`
	// 人群主键id
	CrowdId int64 `json:"crowd_id,omitempty" xml:"crowd_id,omitempty"`
	// 种子人群label信息
	Label *ExtendSeedLabelVo `json:"label,omitempty" xml:"label,omitempty"`
}

var poolExtendSeedCrowdVo = sync.Pool{
	New: func() any {
		return new(ExtendSeedCrowdVo)
	},
}

// GetExtendSeedCrowdVo() 从对象池中获取ExtendSeedCrowdVo
func GetExtendSeedCrowdVo() *ExtendSeedCrowdVo {
	return poolExtendSeedCrowdVo.Get().(*ExtendSeedCrowdVo)
}

// ReleaseExtendSeedCrowdVo 释放ExtendSeedCrowdVo
func ReleaseExtendSeedCrowdVo(v *ExtendSeedCrowdVo) {
	v.CrowdName = ""
	v.CrowdId = 0
	v.Label = nil
	poolExtendSeedCrowdVo.Put(v)
}
