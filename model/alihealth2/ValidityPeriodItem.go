package alihealth2

// ValidityPeriodItem 结构体
type ValidityPeriodItem struct {
	// 请求子明细
	SubItems []ValidityPeriodSubItem `json:"sub_items,omitempty" xml:"sub_items>validity_period_sub_item,omitempty"`
	// 货品名称
	ScItemName string `json:"sc_item_name,omitempty" xml:"sc_item_name,omitempty"`
	// 规格
	Specification string `json:"specification,omitempty" xml:"specification,omitempty"`
	// 批准文号
	ApprovalNumber string `json:"approval_number,omitempty" xml:"approval_number,omitempty"`
	// 条码
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// 生产企业
	Manufacturer string `json:"manufacturer,omitempty" xml:"manufacturer,omitempty"`
	// 货品ID
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
}
