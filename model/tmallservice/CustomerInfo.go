package tmallservice

import (
	"sync"
)

// CustomerInfo 结构体
type CustomerInfo struct {
	// 寄件人名称
	AccountNick string `json:"account_nick,omitempty" xml:"account_nick,omitempty"`
	// 寄件人手机号码
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 寄件地址
	FullAddress string `json:"full_address,omitempty" xml:"full_address,omitempty"`
}

var poolCustomerInfo = sync.Pool{
	New: func() any {
		return new(CustomerInfo)
	},
}

// GetCustomerInfo() 从对象池中获取CustomerInfo
func GetCustomerInfo() *CustomerInfo {
	return poolCustomerInfo.Get().(*CustomerInfo)
}

// ReleaseCustomerInfo 释放CustomerInfo
func ReleaseCustomerInfo(v *CustomerInfo) {
	v.AccountNick = ""
	v.Phone = ""
	v.FullAddress = ""
	poolCustomerInfo.Put(v)
}
