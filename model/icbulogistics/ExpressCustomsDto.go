package icbulogistics

// ExpressCustomsDto 结构体
type ExpressCustomsDto struct {
	// 申报金额
	DeclarationAmount string `json:"declaration_amount,omitempty" xml:"declaration_amount,omitempty"`
	// 是否正式报关
	NeedCustomsClearance string `json:"need_customs_clearance,omitempty" xml:"need_customs_clearance,omitempty"`
	// 报关币种，出口发货中心默认USD
	DeclarationCurrency string `json:"declaration_currency,omitempty" xml:"declaration_currency,omitempty"`
}
