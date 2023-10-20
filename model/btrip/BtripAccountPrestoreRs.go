package btrip

import (
	"sync"
)

// BtripAccountPrestoreRs 结构体
type BtripAccountPrestoreRs struct {
	// 错误信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 账户信息返回
	Module *BtripAccountPrestoreRs `json:"module,omitempty" xml:"module,omitempty"`
	// 错误编码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 预存账户
	BtripAccountPrestoreRs *BtripAccountDetailRs `json:"btrip_account_prestore_rs,omitempty" xml:"btrip_account_prestore_rs,omitempty"`
	// 月结
	BtripCrediAmountRs *BtripCrediAmountRs `json:"btrip_credi_amount_rs,omitempty" xml:"btrip_credi_amount_rs,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolBtripAccountPrestoreRs = sync.Pool{
	New: func() any {
		return new(BtripAccountPrestoreRs)
	},
}

// GetBtripAccountPrestoreRs() 从对象池中获取BtripAccountPrestoreRs
func GetBtripAccountPrestoreRs() *BtripAccountPrestoreRs {
	return poolBtripAccountPrestoreRs.Get().(*BtripAccountPrestoreRs)
}

// ReleaseBtripAccountPrestoreRs 释放BtripAccountPrestoreRs
func ReleaseBtripAccountPrestoreRs(v *BtripAccountPrestoreRs) {
	v.ResultMsg = ""
	v.Module = nil
	v.ResultCode = 0
	v.BtripAccountPrestoreRs = nil
	v.BtripCrediAmountRs = nil
	v.Success = false
	poolBtripAccountPrestoreRs.Put(v)
}
