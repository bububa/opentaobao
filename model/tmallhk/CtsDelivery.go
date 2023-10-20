package tmallhk

import (
	"sync"
)

// CtsDelivery 结构体
type CtsDelivery struct {
	// 快递公司名称
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 快递单号
	DeliveryNo string `json:"delivery_no,omitempty" xml:"delivery_no,omitempty"`
	// 发件时间
	DeliveryTime string `json:"delivery_time,omitempty" xml:"delivery_time,omitempty"`
}

var poolCtsDelivery = sync.Pool{
	New: func() any {
		return new(CtsDelivery)
	},
}

// GetCtsDelivery() 从对象池中获取CtsDelivery
func GetCtsDelivery() *CtsDelivery {
	return poolCtsDelivery.Get().(*CtsDelivery)
}

// ReleaseCtsDelivery 释放CtsDelivery
func ReleaseCtsDelivery(v *CtsDelivery) {
	v.CompanyName = ""
	v.DeliveryNo = ""
	v.DeliveryTime = ""
	poolCtsDelivery.Put(v)
}
