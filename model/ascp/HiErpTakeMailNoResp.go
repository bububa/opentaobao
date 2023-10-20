package ascp

import (
	"sync"
)

// HiErpTakeMailNoResp 结构体
type HiErpTakeMailNoResp struct {
	// 包裹模型
	PackageInfoList []Integer `json:"package_info_list,omitempty" xml:"package_info_list>integer,omitempty"`
	// scp单号
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 外部订单号
	OutOrderCode string `json:"out_order_code,omitempty" xml:"out_order_code,omitempty"`
	// 包裹数量
	PackageQuantity int64 `json:"package_quantity,omitempty" xml:"package_quantity,omitempty"`
}

var poolHiErpTakeMailNoResp = sync.Pool{
	New: func() any {
		return new(HiErpTakeMailNoResp)
	},
}

// GetHiErpTakeMailNoResp() 从对象池中获取HiErpTakeMailNoResp
func GetHiErpTakeMailNoResp() *HiErpTakeMailNoResp {
	return poolHiErpTakeMailNoResp.Get().(*HiErpTakeMailNoResp)
}

// ReleaseHiErpTakeMailNoResp 释放HiErpTakeMailNoResp
func ReleaseHiErpTakeMailNoResp(v *HiErpTakeMailNoResp) {
	v.PackageInfoList = v.PackageInfoList[:0]
	v.OrderCode = ""
	v.OutOrderCode = ""
	v.PackageQuantity = 0
	poolHiErpTakeMailNoResp.Put(v)
}
