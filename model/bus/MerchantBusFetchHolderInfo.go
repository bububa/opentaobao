package bus

import (
	"sync"
)

// MerchantBusFetchHolderInfo 结构体
type MerchantBusFetchHolderInfo struct {
	// 取票电话
	FetchPhone string `json:"fetch_phone,omitempty" xml:"fetch_phone,omitempty"`
	// 取票人信息
	FetchTicketName string `json:"fetch_ticket_name,omitempty" xml:"fetch_ticket_name,omitempty"`
}

var poolMerchantBusFetchHolderInfo = sync.Pool{
	New: func() any {
		return new(MerchantBusFetchHolderInfo)
	},
}

// GetMerchantBusFetchHolderInfo() 从对象池中获取MerchantBusFetchHolderInfo
func GetMerchantBusFetchHolderInfo() *MerchantBusFetchHolderInfo {
	return poolMerchantBusFetchHolderInfo.Get().(*MerchantBusFetchHolderInfo)
}

// ReleaseMerchantBusFetchHolderInfo 释放MerchantBusFetchHolderInfo
func ReleaseMerchantBusFetchHolderInfo(v *MerchantBusFetchHolderInfo) {
	v.FetchPhone = ""
	v.FetchTicketName = ""
	poolMerchantBusFetchHolderInfo.Put(v)
}
