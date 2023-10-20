package tmallservice

import (
	"sync"
)

// SettleAdjustmentResp 结构体
type SettleAdjustmentResp struct {
	// 结算调整单ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolSettleAdjustmentResp = sync.Pool{
	New: func() any {
		return new(SettleAdjustmentResp)
	},
}

// GetSettleAdjustmentResp() 从对象池中获取SettleAdjustmentResp
func GetSettleAdjustmentResp() *SettleAdjustmentResp {
	return poolSettleAdjustmentResp.Get().(*SettleAdjustmentResp)
}

// ReleaseSettleAdjustmentResp 释放SettleAdjustmentResp
func ReleaseSettleAdjustmentResp(v *SettleAdjustmentResp) {
	v.Id = 0
	poolSettleAdjustmentResp.Put(v)
}
