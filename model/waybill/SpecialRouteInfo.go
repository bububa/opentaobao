package waybill

import (
	"sync"
)

// SpecialRouteInfo 结构体
type SpecialRouteInfo struct {
	// 快递公司code
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 到货区域
	ReceiveArea *AddressArea `json:"receive_area,omitempty" xml:"receive_area,omitempty"`
}

var poolSpecialRouteInfo = sync.Pool{
	New: func() any {
		return new(SpecialRouteInfo)
	},
}

// GetSpecialRouteInfo() 从对象池中获取SpecialRouteInfo
func GetSpecialRouteInfo() *SpecialRouteInfo {
	return poolSpecialRouteInfo.Get().(*SpecialRouteInfo)
}

// ReleaseSpecialRouteInfo 释放SpecialRouteInfo
func ReleaseSpecialRouteInfo(v *SpecialRouteInfo) {
	v.CpCode = ""
	v.ReceiveArea = nil
	poolSpecialRouteInfo.Put(v)
}
