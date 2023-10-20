package simba

import (
	"sync"
)

// ExtendSeedCrowdRefVo 结构体
type ExtendSeedCrowdRefVo struct {
	// 种子人群信息
	Crowd *ExtendSeedCrowdVo `json:"crowd,omitempty" xml:"crowd,omitempty"`
}

var poolExtendSeedCrowdRefVo = sync.Pool{
	New: func() any {
		return new(ExtendSeedCrowdRefVo)
	},
}

// GetExtendSeedCrowdRefVo() 从对象池中获取ExtendSeedCrowdRefVo
func GetExtendSeedCrowdRefVo() *ExtendSeedCrowdRefVo {
	return poolExtendSeedCrowdRefVo.Get().(*ExtendSeedCrowdRefVo)
}

// ReleaseExtendSeedCrowdRefVo 释放ExtendSeedCrowdRefVo
func ReleaseExtendSeedCrowdRefVo(v *ExtendSeedCrowdRefVo) {
	v.Crowd = nil
	poolExtendSeedCrowdRefVo.Put(v)
}
