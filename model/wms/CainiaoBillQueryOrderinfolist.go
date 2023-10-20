package wms

import (
	"sync"
)

// CainiaoBillQueryOrderinfolist 结构体
type CainiaoBillQueryOrderinfolist struct {
	// 订单信息
	OrderInfo *CainiaoBillQueryOrderinfo `json:"order_info,omitempty" xml:"order_info,omitempty"`
}

var poolCainiaoBillQueryOrderinfolist = sync.Pool{
	New: func() any {
		return new(CainiaoBillQueryOrderinfolist)
	},
}

// GetCainiaoBillQueryOrderinfolist() 从对象池中获取CainiaoBillQueryOrderinfolist
func GetCainiaoBillQueryOrderinfolist() *CainiaoBillQueryOrderinfolist {
	return poolCainiaoBillQueryOrderinfolist.Get().(*CainiaoBillQueryOrderinfolist)
}

// ReleaseCainiaoBillQueryOrderinfolist 释放CainiaoBillQueryOrderinfolist
func ReleaseCainiaoBillQueryOrderinfolist(v *CainiaoBillQueryOrderinfolist) {
	v.OrderInfo = nil
	poolCainiaoBillQueryOrderinfolist.Put(v)
}
