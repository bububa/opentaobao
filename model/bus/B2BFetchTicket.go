package bus

import (
	"sync"
)

// B2BFetchTicket 结构体
type B2BFetchTicket struct {
	// 取票短信
	FetchSms string `json:"fetch_sms,omitempty" xml:"fetch_sms,omitempty"`
	// 取票地址
	FetchTicketAddress string `json:"fetch_ticket_address,omitempty" xml:"fetch_ticket_address,omitempty"`
	// 取票号
	FetchTicketNumber string `json:"fetch_ticket_number,omitempty" xml:"fetch_ticket_number,omitempty"`
	// 取票密码
	FetchTicketPwd string `json:"fetch_ticket_pwd,omitempty" xml:"fetch_ticket_pwd,omitempty"`
}

var poolB2BFetchTicket = sync.Pool{
	New: func() any {
		return new(B2BFetchTicket)
	},
}

// GetB2BFetchTicket() 从对象池中获取B2BFetchTicket
func GetB2BFetchTicket() *B2BFetchTicket {
	return poolB2BFetchTicket.Get().(*B2BFetchTicket)
}

// ReleaseB2BFetchTicket 释放B2BFetchTicket
func ReleaseB2BFetchTicket(v *B2BFetchTicket) {
	v.FetchSms = ""
	v.FetchTicketAddress = ""
	v.FetchTicketNumber = ""
	v.FetchTicketPwd = ""
	poolB2BFetchTicket.Put(v)
}
