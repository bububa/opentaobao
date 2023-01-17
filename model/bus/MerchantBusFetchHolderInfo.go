package bus

// MerchantBusFetchHolderInfo 结构体
type MerchantBusFetchHolderInfo struct {
	// 取票电话
	FetchPhone string `json:"fetch_phone,omitempty" xml:"fetch_phone,omitempty"`
	// 取票人信息
	FetchTicketName string `json:"fetch_ticket_name,omitempty" xml:"fetch_ticket_name,omitempty"`
}
