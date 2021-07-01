package damai

// TicketItemIdOpenParam 结构体
type TicketItemIdOpenParam struct {
	// 票品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 场次id
	PerformId int64 `json:"perform_id,omitempty" xml:"perform_id,omitempty"`
	// 系统id
	SystemId int64 `json:"system_id,omitempty" xml:"system_id,omitempty"`
	// 商户密钥
	SupplierSecret string `json:"supplier_secret,omitempty" xml:"supplier_secret,omitempty"`
}
