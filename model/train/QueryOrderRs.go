package train

// QueryOrderRs 结构体
type QueryOrderRs struct {
	// 票详情
	TicketInfos []TicketInfoDto `json:"ticket_infos,omitempty" xml:"ticket_infos>ticket_info_dto,omitempty"`
	// 子单详情
	TapSubOrders []TapSubOrderVO `json:"tap_sub_orders,omitempty" xml:"tap_sub_orders>tap_sub_order_vo,omitempty"`
	// 主单详情
	TapOrder *TapOrderVO `json:"tap_order,omitempty" xml:"tap_order,omitempty"`
}
