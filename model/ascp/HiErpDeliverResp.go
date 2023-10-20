package ascp

import (
	"sync"
)

// HiErpDeliverResp 结构体
type HiErpDeliverResp struct {
	// SCP单号
	ScpCode string `json:"scp_code,omitempty" xml:"scp_code,omitempty"`
	// 外部订单号
	OutOrderCode string `json:"out_order_code,omitempty" xml:"out_order_code,omitempty"`
}

var poolHiErpDeliverResp = sync.Pool{
	New: func() any {
		return new(HiErpDeliverResp)
	},
}

// GetHiErpDeliverResp() 从对象池中获取HiErpDeliverResp
func GetHiErpDeliverResp() *HiErpDeliverResp {
	return poolHiErpDeliverResp.Get().(*HiErpDeliverResp)
}

// ReleaseHiErpDeliverResp 释放HiErpDeliverResp
func ReleaseHiErpDeliverResp(v *HiErpDeliverResp) {
	v.ScpCode = ""
	v.OutOrderCode = ""
	poolHiErpDeliverResp.Put(v)
}
