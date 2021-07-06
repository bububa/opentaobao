package flight

// TicketingIssueRequestDto 结构体
type TicketingIssueRequestDto struct {
	// 出票信息
	IssueList []TicketingPsgItemDto `json:"issue_list,omitempty" xml:"issue_list>ticketing_psg_item_dto,omitempty"`
	// 飞猪订单号
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 国内国际标识
	DomesticIntl int64 `json:"domestic_intl,omitempty" xml:"domestic_intl,omitempty"`
}
