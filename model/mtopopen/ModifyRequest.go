package mtopopen

import (
	"sync"
)

// ModifyRequest 结构体
type ModifyRequest struct {
	// 订单ID
	OrderCode string `json:"order_code,omitempty" xml:"order_code,omitempty"`
	// 快递公司标准编码
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 快递单号
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 修改后收件人
	Fetcher string `json:"fetcher,omitempty" xml:"fetcher,omitempty"`
	// 修改后派送时间
	DeliveryTime string `json:"delivery_time,omitempty" xml:"delivery_time,omitempty"`
}

var poolModifyRequest = sync.Pool{
	New: func() any {
		return new(ModifyRequest)
	},
}

// GetModifyRequest() 从对象池中获取ModifyRequest
func GetModifyRequest() *ModifyRequest {
	return poolModifyRequest.Get().(*ModifyRequest)
}

// ReleaseModifyRequest 释放ModifyRequest
func ReleaseModifyRequest(v *ModifyRequest) {
	v.OrderCode = ""
	v.CpCode = ""
	v.MailNo = ""
	v.Fetcher = ""
	v.DeliveryTime = ""
	poolModifyRequest.Put(v)
}
