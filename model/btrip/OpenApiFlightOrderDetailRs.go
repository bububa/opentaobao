package btrip

// OpenApiFlightOrderDetailRs 结构体
type OpenApiFlightOrderDetailRs struct {
	// 订单费用列表
	PriceInfoList []PriceInfo `json:"price_info_list,omitempty" xml:"price_info_list>price_info,omitempty"`
	// 退票列表
	FlightRefundTicketInfoList []FlightRefundTicketInfo `json:"flight_refund_ticket_info_list,omitempty" xml:"flight_refund_ticket_info_list>flight_refund_ticket_info,omitempty"`
	// 航班信息
	FlightInfoList []FlightInfo `json:"flight_info_list,omitempty" xml:"flight_info_list>flight_info,omitempty"`
	// 保险列表
	InsuranceInfoList []InsuranceInfo `json:"insurance_info_list,omitempty" xml:"insurance_info_list>insurance_info,omitempty"`
	// 正票列表
	FlightTicketInfoList []FlightTicketInfo `json:"flight_ticket_info_list,omitempty" xml:"flight_ticket_info_list>flight_ticket_info,omitempty"`
	// 出行人列表
	PassengerInfoList []PassengerInfo `json:"passenger_info_list,omitempty" xml:"passenger_info_list>passenger_info,omitempty"`
	// 改签票列表
	FlightChangeTicketInfoList []FlightChangeTicketInfo `json:"flight_change_ticket_info_list,omitempty" xml:"flight_change_ticket_info_list>flight_change_ticket_info,omitempty"`
	// 订单基础信息
	OrderBaseInfo *OrderBaseInfo `json:"order_base_info,omitempty" xml:"order_base_info,omitempty"`
	// 发票信息
	InvoiceInfo *InvoiceInfo `json:"invoice_info,omitempty" xml:"invoice_info,omitempty"`
}
