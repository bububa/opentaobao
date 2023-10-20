package bus

import (
	"sync"
)

// B2BFetchHolderInfo 结构体
type B2BFetchHolderInfo struct {
	// 取票人证件号码
	FetchCertNumber string `json:"fetch_cert_number,omitempty" xml:"fetch_cert_number,omitempty"`
	// 取票人证件类型
	FetchCertType string `json:"fetch_cert_type,omitempty" xml:"fetch_cert_type,omitempty"`
	// 取票人电话
	FetchPhone string `json:"fetch_phone,omitempty" xml:"fetch_phone,omitempty"`
	// 取票人姓名
	FetchTicketName string `json:"fetch_ticket_name,omitempty" xml:"fetch_ticket_name,omitempty"`
}

var poolB2BFetchHolderInfo = sync.Pool{
	New: func() any {
		return new(B2BFetchHolderInfo)
	},
}

// GetB2BFetchHolderInfo() 从对象池中获取B2BFetchHolderInfo
func GetB2BFetchHolderInfo() *B2BFetchHolderInfo {
	return poolB2BFetchHolderInfo.Get().(*B2BFetchHolderInfo)
}

// ReleaseB2BFetchHolderInfo 释放B2BFetchHolderInfo
func ReleaseB2BFetchHolderInfo(v *B2BFetchHolderInfo) {
	v.FetchCertNumber = ""
	v.FetchCertType = ""
	v.FetchPhone = ""
	v.FetchTicketName = ""
	poolB2BFetchHolderInfo.Put(v)
}
