package einvoice

// InvoiceItem 结构体
type InvoiceItem struct {
	// 发票项目名称（或商品名称）
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 价税合计。(等于sumPrice和tax之和)
	Amount string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 发票行性质。0表示正常行，1表示折扣行，2表示被折扣行。比如充电器单价100元，折扣10元，则明细为2行，充电器行性质为2，折扣行性质为1。如果充电器没有折扣，则值应为0
	RowType string `json:"row_type,omitempty" xml:"row_type,omitempty"`
	// 规格型号,可选
	Specification string `json:"specification,omitempty" xml:"specification,omitempty"`
	// 总价，格式：100.00(不含税)
	SumPrice string `json:"sum_price,omitempty" xml:"sum_price,omitempty"`
	// 税额
	Tax string `json:"tax,omitempty" xml:"tax,omitempty"`
	// 单价，格式：100.00(不含税)
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 数量
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 税率。税率只能为0或0.03或0.04或0.06或0.11或0.13或0.17
	TaxRate string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	// 单位
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 发票项目编号（或商品编号）
	ItemNo string `json:"item_no,omitempty" xml:"item_no,omitempty"`
	// 淘宝子订单号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 零税率标识，0=出口零税率，1=免税，2=不征收，3=普通零税率
	ZeroRateFlag string `json:"zero_rate_flag,omitempty" xml:"zero_rate_flag,omitempty"`
	// 商品的外部系统id，如果有sku则取sku的outerId，否则取item的outerId，，阿里发票平台自动开票时才有
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 商品IMIE号(不用传，将废弃)
	Imei string `json:"imei,omitempty" xml:"imei,omitempty"`
	// 是否运费行标识，true:运费行，false:非运费行
	IsPostFeeRow bool `json:"is_post_fee_row,omitempty" xml:"is_post_fee_row,omitempty"`
}
