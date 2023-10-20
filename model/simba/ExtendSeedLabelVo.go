package simba

import (
	"sync"
)

// ExtendSeedLabelVo 结构体
type ExtendSeedLabelVo struct {
}

var poolExtendSeedLabelVo = sync.Pool{
	New: func() any {
		return new(ExtendSeedLabelVo)
	},
}

// GetExtendSeedLabelVo() 从对象池中获取ExtendSeedLabelVo
func GetExtendSeedLabelVo() *ExtendSeedLabelVo {
	return poolExtendSeedLabelVo.Get().(*ExtendSeedLabelVo)
}

// ReleaseExtendSeedLabelVo 释放ExtendSeedLabelVo
func ReleaseExtendSeedLabelVo(v *ExtendSeedLabelVo) {
	poolExtendSeedLabelVo.Put(v)
}
