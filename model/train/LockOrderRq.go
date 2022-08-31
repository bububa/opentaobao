package train

// LockOrderRq 结构体
type LockOrderRq struct {
	// 票信息
	Tickets []TicketDto `json:"tickets,omitempty" xml:"tickets>ticket_dto,omitempty"`
	// 主单号
	TtpOrderId int64 `json:"ttp_order_id,omitempty" xml:"ttp_order_id,omitempty"`
}
