package wdk

import (
	"sync"
)

// Customer 结构体
type Customer struct {
	// 收货人地址
	BuyerAddress string `json:"buyer_address,omitempty" xml:"buyer_address,omitempty"`
	// 收货人联系电话
	BuyerPhone string `json:"buyer_phone,omitempty" xml:"buyer_phone,omitempty"`
	// 收货人姓名
	BuyerName string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
}

var poolCustomer = sync.Pool{
	New: func() any {
		return new(Customer)
	},
}

// GetCustomer() 从对象池中获取Customer
func GetCustomer() *Customer {
	return poolCustomer.Get().(*Customer)
}

// ReleaseCustomer 释放Customer
func ReleaseCustomer(v *Customer) {
	v.BuyerAddress = ""
	v.BuyerPhone = ""
	v.BuyerName = ""
	poolCustomer.Put(v)
}
