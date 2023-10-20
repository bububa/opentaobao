package btrip

import (
	"sync"
)

// OrderTicketInfo 结构体
type OrderTicketInfo struct {
	// pnr编码
	PnrCode string `json:"pnr_code,omitempty" xml:"pnr_code,omitempty"`
	// 票号
	TicketNo string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	// 票状态
	TicketStatus string `json:"ticket_status,omitempty" xml:"ticket_status,omitempty"`
	// openTicket状态
	OpenTicketStatus string `json:"open_ticket_status,omitempty" xml:"open_ticket_status,omitempty"`
}

var poolOrderTicketInfo = sync.Pool{
	New: func() any {
		return new(OrderTicketInfo)
	},
}

// GetOrderTicketInfo() 从对象池中获取OrderTicketInfo
func GetOrderTicketInfo() *OrderTicketInfo {
	return poolOrderTicketInfo.Get().(*OrderTicketInfo)
}

// ReleaseOrderTicketInfo 释放OrderTicketInfo
func ReleaseOrderTicketInfo(v *OrderTicketInfo) {
	v.PnrCode = ""
	v.TicketNo = ""
	v.TicketStatus = ""
	v.OpenTicketStatus = ""
	poolOrderTicketInfo.Put(v)
}
