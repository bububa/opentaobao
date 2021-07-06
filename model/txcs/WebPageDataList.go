package txcs

// WebPageDataList 结构体
type WebPageDataList struct {
	// 对账单号
	StatementBillCode string `json:"statement_bill_code,omitempty" xml:"statement_bill_code,omitempty"`
	// 店仓编码
	ShopCode string `json:"shop_code,omitempty" xml:"shop_code,omitempty"`
	// 商品数量
	ItemQuantity string `json:"item_quantity,omitempty" xml:"item_quantity,omitempty"`
	// 来源单据描述
	SourceTypeDesc string `json:"source_type_desc,omitempty" xml:"source_type_desc,omitempty"`
	// 业务单据编码
	BizBillCode string `json:"biz_bill_code,omitempty" xml:"biz_bill_code,omitempty"`
	// 结算方式描述
	SettleWayDesc string `json:"settle_way_desc,omitempty" xml:"settle_way_desc,omitempty"`
	// 店仓名称
	ShopName string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
	// 业务订单编号
	BizOrderCode string `json:"biz_order_code,omitempty" xml:"biz_order_code,omitempty"`
	// 合同版本号
	ContractVersion string `json:"contract_version,omitempty" xml:"contract_version,omitempty"`
	// 税率
	TaxRate string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	// 单位
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 来源单据类型
	SourceType string `json:"source_type,omitempty" xml:"source_type,omitempty"`
	// 费用编码
	FeeCode string `json:"fee_code,omitempty" xml:"fee_code,omitempty"`
	// 费用名称
	FeeName string `json:"fee_name,omitempty" xml:"fee_name,omitempty"`
	// 未税金额
	UntaxAmount string `json:"untax_amount,omitempty" xml:"untax_amount,omitempty"`
	// 结算金额
	SettlementAmount string `json:"settlement_amount,omitempty" xml:"settlement_amount,omitempty"`
	// 税额
	TaxAmount string `json:"tax_amount,omitempty" xml:"tax_amount,omitempty"`
	// 合同编码
	ContractCode string `json:"contract_code,omitempty" xml:"contract_code,omitempty"`
	// 业务编码
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 扩展信息
	BizExtAttr *BizExtAttr `json:"biz_ext_attr,omitempty" xml:"biz_ext_attr,omitempty"`
	// 叶子类目编码
	LeafCatId int64 `json:"leaf_cat_id,omitempty" xml:"leaf_cat_id,omitempty"`
	// 业务时间
	BizTime int64 `json:"biz_time,omitempty" xml:"biz_time,omitempty"`
	// 币种
	Currency *Currency `json:"currency,omitempty" xml:"currency,omitempty"`
	// 结算方式描述
	SettleWay int64 `json:"settle_way,omitempty" xml:"settle_way,omitempty"`
}
