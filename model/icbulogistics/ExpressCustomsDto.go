package icbulogistics

// ExpressCustomsDto 结构体
type ExpressCustomsDto struct {
	// 申报金额
	DeclarationAmount string `json:"declaration_amount,omitempty" xml:"declaration_amount,omitempty"`
	// 是否正式报关
	NeedCustomsClearance string `json:"need_customs_clearance,omitempty" xml:"need_customs_clearance,omitempty"`
	// 报关币种，出口发货中心默认USD
	DeclarationCurrency string `json:"declaration_currency,omitempty" xml:"declaration_currency,omitempty"`
	// 增值税类型，枚举取值：VAT、IOSS、VOEC
	VatType string `json:"vat_type,omitempty" xml:"vat_type,omitempty"`
	// 增值税税号
	VatNumber string `json:"vat_number,omitempty" xml:"vat_number,omitempty"`
	// 纳税人识别号
	TaxpayerId string `json:"taxpayer_id,omitempty" xml:"taxpayer_id,omitempty"`
	// 欧盟EORI
	EoriNumber string `json:"eori_number,omitempty" xml:"eori_number,omitempty"`
}
