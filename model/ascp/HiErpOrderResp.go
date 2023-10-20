package ascp

import (
	"sync"
)

// HiErpOrderResp 结构体
type HiErpOrderResp struct {
	// 外部订单号
	OutOrderCode string `json:"out_order_code,omitempty" xml:"out_order_code,omitempty"`
	// SCP单号，履约单号
	ScpOrderCode string `json:"scp_order_code,omitempty" xml:"scp_order_code,omitempty"`
	// 扩展字段，json格式
	Feature string `json:"feature,omitempty" xml:"feature,omitempty"`
}

var poolHiErpOrderResp = sync.Pool{
	New: func() any {
		return new(HiErpOrderResp)
	},
}

// GetHiErpOrderResp() 从对象池中获取HiErpOrderResp
func GetHiErpOrderResp() *HiErpOrderResp {
	return poolHiErpOrderResp.Get().(*HiErpOrderResp)
}

// ReleaseHiErpOrderResp 释放HiErpOrderResp
func ReleaseHiErpOrderResp(v *HiErpOrderResp) {
	v.OutOrderCode = ""
	v.ScpOrderCode = ""
	v.Feature = ""
	poolHiErpOrderResp.Put(v)
}
