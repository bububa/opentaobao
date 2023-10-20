package ascpchannel

import (
	"sync"
)

// ExtOrderReceiverRequest 结构体
type ExtOrderReceiverRequest struct {
	// 收货人名称
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 收货人地址
	DetailAddress string `json:"detail_address,omitempty" xml:"detail_address,omitempty"`
	// 地址编码
	DivisionId int64 `json:"division_id,omitempty" xml:"division_id,omitempty"`
}

var poolExtOrderReceiverRequest = sync.Pool{
	New: func() any {
		return new(ExtOrderReceiverRequest)
	},
}

// GetExtOrderReceiverRequest() 从对象池中获取ExtOrderReceiverRequest
func GetExtOrderReceiverRequest() *ExtOrderReceiverRequest {
	return poolExtOrderReceiverRequest.Get().(*ExtOrderReceiverRequest)
}

// ReleaseExtOrderReceiverRequest 释放ExtOrderReceiverRequest
func ReleaseExtOrderReceiverRequest(v *ExtOrderReceiverRequest) {
	v.ContactName = ""
	v.DetailAddress = ""
	v.DivisionId = 0
	poolExtOrderReceiverRequest.Put(v)
}
