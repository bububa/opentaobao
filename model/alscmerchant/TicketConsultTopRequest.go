package alscmerchant

import (
	"sync"
)

// TicketConsultTopRequest 结构体
type TicketConsultTopRequest struct {
	// 请求id
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 本地门店id，长度32位
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 券码code
	TicketCode string `json:"ticket_code,omitempty" xml:"ticket_code,omitempty"`
}

var poolTicketConsultTopRequest = sync.Pool{
	New: func() any {
		return new(TicketConsultTopRequest)
	},
}

// GetTicketConsultTopRequest() 从对象池中获取TicketConsultTopRequest
func GetTicketConsultTopRequest() *TicketConsultTopRequest {
	return poolTicketConsultTopRequest.Get().(*TicketConsultTopRequest)
}

// ReleaseTicketConsultTopRequest 释放TicketConsultTopRequest
func ReleaseTicketConsultTopRequest(v *TicketConsultTopRequest) {
	v.RequestId = ""
	v.ShopId = ""
	v.TicketCode = ""
	poolTicketConsultTopRequest.Put(v)
}
