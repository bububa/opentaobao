package pur

// DeliveryItemTopDto 结构体
type DeliveryItemTopDto struct {
	// 收货地址
	AddressInfo string `json:"address_info,omitempty" xml:"address_info,omitempty"`
	// 订单编号
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 饿了么行id
	BizId string `json:"biz_id,omitempty" xml:"biz_id,omitempty"`
	// 采购类目
	CategoryCode string `json:"category_code,omitempty" xml:"category_code,omitempty"`
	// 用途code
	CategoryUse string `json:"category_use,omitempty" xml:"category_use,omitempty"`
	// 公司code
	CompanyCode string `json:"company_code,omitempty" xml:"company_code,omitempty"`
	// 币种
	CurrencyCode string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
	// 发运单位
	DeliveryUnit string `json:"delivery_unit,omitempty" xml:"delivery_unit,omitempty"`
	// 实际需求方
	Demander string `json:"demander,omitempty" xml:"demander,omitempty"`
	// 费用归属期间
	ExpenseMonth string `json:"expense_month,omitempty" xml:"expense_month,omitempty"`
	// 费用归属结束时间
	ExpenseMonthEnd string `json:"expense_month_end,omitempty" xml:"expense_month_end,omitempty"`
	// 费用归属开始时间
	ExpenseMonthStart string `json:"expense_month_start,omitempty" xml:"expense_month_start,omitempty"`
	// 拓展字段
	ExtStr string `json:"ext_str,omitempty" xml:"ext_str,omitempty"`
	// 需求说明
	ItemDescription string `json:"item_description,omitempty" xml:"item_description,omitempty"`
	// 订单行id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 物品名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 购买方式,如按数量,按金额
	OrderType string `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// 联系方式
	PhoneNo string `json:"phone_no,omitempty" xml:"phone_no,omitempty"`
	// 发货数量/金额
	QuantityDelivered string `json:"quantity_delivered,omitempty" xml:"quantity_delivered,omitempty"`
	// 接收单位
	ReceiveUnit string `json:"receive_unit,omitempty" xml:"receive_unit,omitempty"`
	// 接收人
	Receiver string `json:"receiver,omitempty" xml:"receiver,omitempty"`
	// 规格属性
	Specification string `json:"specification,omitempty" xml:"specification,omitempty"`
	// 结算日期（结构化发货才需要）
	SettlementDate string `json:"settlement_date,omitempty" xml:"settlement_date,omitempty"`
	// 税码
	TaxCode string `json:"tax_code,omitempty" xml:"tax_code,omitempty"`
	// 税率
	TaxRate string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	// 单价
	UnitPrice string `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
	// 计价单位
	Uom string `json:"uom,omitempty" xml:"uom,omitempty"`
	// 结构化信息
	StructureItem string `json:"structure_item,omitempty" xml:"structure_item,omitempty"`
	// 组织id
	PurchaseOrgId int64 `json:"purchase_org_id,omitempty" xml:"purchase_org_id,omitempty"`
	// 供应商id
	SupplierId int64 `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 是否按明细
	DeliveryByDetail bool `json:"delivery_by_detail,omitempty" xml:"delivery_by_detail,omitempty"`
}
