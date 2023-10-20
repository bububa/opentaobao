package btrip

import (
	"sync"
)

// Tips 结构体
type Tips struct {
}

var poolTips = sync.Pool{
	New: func() any {
		return new(Tips)
	},
}

// GetTips() 从对象池中获取Tips
func GetTips() *Tips {
	return poolTips.Get().(*Tips)
}

// ReleaseTips 释放Tips
func ReleaseTips(v *Tips) {
	poolTips.Put(v)
}
