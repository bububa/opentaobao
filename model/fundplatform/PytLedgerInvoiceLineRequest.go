package fundplatform

// PytLedgerInvoiceLineRequest 结构体
type PytLedgerInvoiceLineRequest struct {
	// 含税金额
	AmountWithTax string `json:"amount_with_tax,omitempty" xml:"amount_with_tax,omitempty"`
	// 不含税金额
	AmountWithoutTax string `json:"amount_without_tax,omitempty" xml:"amount_without_tax,omitempty"`
	// 扣除额
	Deductions string `json:"deductions,omitempty" xml:"deductions,omitempty"`
	// 税收分类编码
	GoodsTaxNo string `json:"goods_tax_no,omitempty" xml:"goods_tax_no,omitempty"`
	// 编码版本号
	GoodsTaxNoVersion string `json:"goods_tax_no_version,omitempty" xml:"goods_tax_no_version,omitempty"`
	// 数量
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 规格型号
	Specifications string `json:"specifications,omitempty" xml:"specifications,omitempty"`
	// 税额
	TaxAmount string `json:"tax_amount,omitempty" xml:"tax_amount,omitempty"`
	// 是否享受税收优惠政策
	TaxPre string `json:"tax_pre,omitempty" xml:"tax_pre,omitempty"`
	// 享受税收优惠政策内容
	TaxPreCon string `json:"tax_pre_con,omitempty" xml:"tax_pre_con,omitempty"`
	// 税率
	TaxRate string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	// 通行日期止
	TollEndDate string `json:"toll_end_date,omitempty" xml:"toll_end_date,omitempty"`
	// 通行日期起
	TollStartDate string `json:"toll_start_date,omitempty" xml:"toll_start_date,omitempty"`
	// 数量单位
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 单价
	UnitPrice string `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
	// 零税率标志
	ZeroTax string `json:"zero_tax,omitempty" xml:"zero_tax,omitempty"`
	// 项目名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 项目编码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
}
