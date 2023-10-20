package btrip

import (
	"sync"
)

// BtripAccountDetailRs 结构体
type BtripAccountDetailRs struct {
	// 账户状态1:激活 0：冻结
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 账户状态描述。冻结，激活
	StatusDesc string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	// 可用余额，分
	TotalAmount int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// 账户余额，分
	Balance int64 `json:"balance,omitempty" xml:"balance,omitempty"`
	// 冻结金额，分
	FreezeAmount int64 `json:"freeze_amount,omitempty" xml:"freeze_amount,omitempty"`
}

var poolBtripAccountDetailRs = sync.Pool{
	New: func() any {
		return new(BtripAccountDetailRs)
	},
}

// GetBtripAccountDetailRs() 从对象池中获取BtripAccountDetailRs
func GetBtripAccountDetailRs() *BtripAccountDetailRs {
	return poolBtripAccountDetailRs.Get().(*BtripAccountDetailRs)
}

// ReleaseBtripAccountDetailRs 释放BtripAccountDetailRs
func ReleaseBtripAccountDetailRs(v *BtripAccountDetailRs) {
	v.Status = ""
	v.StatusDesc = ""
	v.TotalAmount = 0
	v.Balance = 0
	v.FreezeAmount = 0
	poolBtripAccountDetailRs.Put(v)
}
