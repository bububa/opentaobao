package btrip

import (
	"sync"
)

// CabinRs 结构体
type CabinRs struct {
}

var poolCabinRs = sync.Pool{
	New: func() any {
		return new(CabinRs)
	},
}

// GetCabinRs() 从对象池中获取CabinRs
func GetCabinRs() *CabinRs {
	return poolCabinRs.Get().(*CabinRs)
}

// ReleaseCabinRs 释放CabinRs
func ReleaseCabinRs(v *CabinRs) {
	poolCabinRs.Put(v)
}
