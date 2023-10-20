package drugtrace

import (
	"sync"
)

// BillIdentity 结构体
type BillIdentity struct {
	// 单据类型
	BillType string `json:"bill_type,omitempty" xml:"bill_type,omitempty"`
	// 单据iD
	BillId int64 `json:"bill_id,omitempty" xml:"bill_id,omitempty"`
}

var poolBillIdentity = sync.Pool{
	New: func() any {
		return new(BillIdentity)
	},
}

// GetBillIdentity() 从对象池中获取BillIdentity
func GetBillIdentity() *BillIdentity {
	return poolBillIdentity.Get().(*BillIdentity)
}

// ReleaseBillIdentity 释放BillIdentity
func ReleaseBillIdentity(v *BillIdentity) {
	v.BillType = ""
	v.BillId = 0
	poolBillIdentity.Put(v)
}
