package wdk

// WdkOpenOrderFinanceBillDo 结构体
type WdkOpenOrderFinanceBillDo struct {
	// alipay：支付宝
	PayChannel string `json:"pay_channel,omitempty" xml:"pay_channel,omitempty"`
	// 盒马订单号
	HmOrderId string `json:"hm_order_id,omitempty" xml:"hm_order_id,omitempty"`
	// app：线上， pos：线下
	OrderChannel string `json:"order_channel,omitempty" xml:"order_channel,omitempty"`
	// 经营店id
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 商家编码
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 淘系订单号
	TpOrderId string `json:"tp_order_id,omitempty" xml:"tp_order_id,omitempty"`
	// 账单日期
	Dt string `json:"dt,omitempty" xml:"dt,omitempty"`
	// 支付宝二级商户ID
	Smid string `json:"smid,omitempty" xml:"smid,omitempty"`
	// 金额，正向交易时数值是正数，退款时数值是负数
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 业务类型。PAY：支付， REFUND：退款
	BizType string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 支付机构支付流水号
	TradeNo string `json:"trade_no,omitempty" xml:"trade_no,omitempty"`
	// 财务组织名称
	FinanceOrganizationName string `json:"finance_organization_name,omitempty" xml:"finance_organization_name,omitempty"`
	// 财务组织code
	FinanceOrganizationCode string `json:"finance_organization_code,omitempty" xml:"finance_organization_code,omitempty"`
}
