package ieagency

import (
	"sync"
)

// IeIssueTicketVo 结构体
type IeIssueTicketVo struct {
	// 乘机人票信息
	PassengerTicketVos []IePassengerTicketVo `json:"passenger_ticket_vos,omitempty" xml:"passenger_ticket_vos>ie_passenger_ticket_vo,omitempty"`
	// 预定订单pnr信息
	UpdatePnrVos []IeBookPnrVo `json:"update_pnr_vos,omitempty" xml:"update_pnr_vos>ie_book_pnr_vo,omitempty"`
	// 订单备注
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
	// 预定订单id
	BookOrderId int64 `json:"book_order_id,omitempty" xml:"book_order_id,omitempty"`
}

var poolIeIssueTicketVo = sync.Pool{
	New: func() any {
		return new(IeIssueTicketVo)
	},
}

// GetIeIssueTicketVo() 从对象池中获取IeIssueTicketVo
func GetIeIssueTicketVo() *IeIssueTicketVo {
	return poolIeIssueTicketVo.Get().(*IeIssueTicketVo)
}

// ReleaseIeIssueTicketVo 释放IeIssueTicketVo
func ReleaseIeIssueTicketVo(v *IeIssueTicketVo) {
	v.PassengerTicketVos = v.PassengerTicketVos[:0]
	v.UpdatePnrVos = v.UpdatePnrVos[:0]
	v.Memo = ""
	v.BookOrderId = 0
	poolIeIssueTicketVo.Put(v)
}
