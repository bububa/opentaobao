package fpm

// RegisterInvoiceLineDto 结构体
type RegisterInvoiceLineDto struct {
	// 发票行类型
	InvoiceLineType string `json:"invoice_line_type,omitempty" xml:"invoice_line_type,omitempty"`
	// 单位
	QuantityUnit string `json:"quantity_unit,omitempty" xml:"quantity_unit,omitempty"`
	// 货物来源国
	GoodsSourceCountry string `json:"goods_source_country,omitempty" xml:"goods_source_country,omitempty"`
	// 税种2
	TaxCategory2 string `json:"tax_category2,omitempty" xml:"tax_category2,omitempty"`
	// 免税注释
	DutyFreeMemo string `json:"duty_free_memo,omitempty" xml:"duty_free_memo,omitempty"`
	// 汇率
	ExchangeRate string `json:"exchange_rate,omitempty" xml:"exchange_rate,omitempty"`
	// 本币税额2
	LocalCurrencyTaxAmount2 string `json:"local_currency_tax_amount2,omitempty" xml:"local_currency_tax_amount2,omitempty"`
	// 本币税额1
	LocalCurrencyTaxAmount1 string `json:"local_currency_tax_amount1,omitempty" xml:"local_currency_tax_amount1,omitempty"`
	// 规格型号
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 币种
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 税种
	TaxCategories string `json:"tax_categories,omitempty" xml:"tax_categories,omitempty"`
	// 税率2
	TaxRate2 string `json:"tax_rate2,omitempty" xml:"tax_rate2,omitempty"`
	// 货物名称
	GoodsDesc string `json:"goods_desc,omitempty" xml:"goods_desc,omitempty"`
	// 供货地点
	PlaceOfSupply string `json:"place_of_supply,omitempty" xml:"place_of_supply,omitempty"`
	// 单价
	UnitPrice string `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
	// 含税金额
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 本币税额
	LocalCurrencyTaxAmount string `json:"local_currency_tax_amount,omitempty" xml:"local_currency_tax_amount,omitempty"`
	// 数量
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 本币不含税金额
	LocalCurrencyExcludingTaxAmount string `json:"local_currency_excluding_tax_amount,omitempty" xml:"local_currency_excluding_tax_amount,omitempty"`
	// 本币含税总金额
	LocalCurrencyAmount string `json:"local_currency_amount,omitempty" xml:"local_currency_amount,omitempty"`
	// 税额1
	TaxAmount1 string `json:"tax_amount1,omitempty" xml:"tax_amount1,omitempty"`
	// 税额2
	TaxAmount2 string `json:"tax_amount2,omitempty" xml:"tax_amount2,omitempty"`
	// 本币币种
	LocalCurrencyCode string `json:"local_currency_code,omitempty" xml:"local_currency_code,omitempty"`
	// 服务日期/供货日期
	SupplyDate string `json:"supply_date,omitempty" xml:"supply_date,omitempty"`
	// 服务核算代码
	ServiceAccountingCode string `json:"service_accounting_code,omitempty" xml:"service_accounting_code,omitempty"`
	// 税率
	TaxRate string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	// 是否海外发票
	ReverseCharge string `json:"reverse_charge,omitempty" xml:"reverse_charge,omitempty"`
	// 本币单价
	LocalCurrencyPrice string `json:"local_currency_price,omitempty" xml:"local_currency_price,omitempty"`
	// 零税率类型
	ZeroRateFlag string `json:"zero_rate_flag,omitempty" xml:"zero_rate_flag,omitempty"`
	// 税额
	TaxAmount string `json:"tax_amount,omitempty" xml:"tax_amount,omitempty"`
	// 不含税金额
	ExcludingTaxAmount string `json:"excluding_tax_amount,omitempty" xml:"excluding_tax_amount,omitempty"`
	// 行号
	RowNo int64 `json:"row_no,omitempty" xml:"row_no,omitempty"`
}
