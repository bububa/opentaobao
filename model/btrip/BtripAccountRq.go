package btrip

import (
	"sync"
)

// BtripAccountRq 结构体
type BtripAccountRq struct {
	// CorpId
	SubChannel string `json:"sub_channel,omitempty" xml:"sub_channel,omitempty"`
}

var poolBtripAccountRq = sync.Pool{
	New: func() any {
		return new(BtripAccountRq)
	},
}

// GetBtripAccountRq() 从对象池中获取BtripAccountRq
func GetBtripAccountRq() *BtripAccountRq {
	return poolBtripAccountRq.Get().(*BtripAccountRq)
}

// ReleaseBtripAccountRq 释放BtripAccountRq
func ReleaseBtripAccountRq(v *BtripAccountRq) {
	v.SubChannel = ""
	poolBtripAccountRq.Put(v)
}
