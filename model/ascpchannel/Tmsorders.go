package ascpchannel

import (
	"sync"
)

// Tmsorders 结构体
type Tmsorders struct {
	// 包裹明细列表
	TmsItems []Tmsitems `json:"tms_items,omitempty" xml:"tms_items>tmsitems,omitempty"`
	// 运单号
	TmsOrderCode string `json:"tms_order_code,omitempty" xml:"tms_order_code,omitempty"`
	// 快递公司code.调用 taobao.logistics.companies.get 获取
	TmsServiceCode string `json:"tms_service_code,omitempty" xml:"tms_service_code,omitempty"`
	// 快递公司名称
	TmsServiceName string `json:"tms_service_name,omitempty" xml:"tms_service_name,omitempty"`
}

var poolTmsorders = sync.Pool{
	New: func() any {
		return new(Tmsorders)
	},
}

// GetTmsorders() 从对象池中获取Tmsorders
func GetTmsorders() *Tmsorders {
	return poolTmsorders.Get().(*Tmsorders)
}

// ReleaseTmsorders 释放Tmsorders
func ReleaseTmsorders(v *Tmsorders) {
	v.TmsItems = v.TmsItems[:0]
	v.TmsOrderCode = ""
	v.TmsServiceCode = ""
	v.TmsServiceName = ""
	poolTmsorders.Put(v)
}
