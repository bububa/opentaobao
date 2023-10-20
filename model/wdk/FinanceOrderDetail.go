package wdk

import (
	"sync"
)

// FinanceOrderDetail 结构体
type FinanceOrderDetail struct {
	// 币种
	Currency string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 税率
	TaxRate string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	// 销售渠道
	SaleChannel string `json:"sale_channel,omitempty" xml:"sale_channel,omitempty"`
	// 销售来源
	SaleSource string `json:"sale_source,omitempty" xml:"sale_source,omitempty"`
	// 交易类型
	TradeType string `json:"trade_type,omitempty" xml:"trade_type,omitempty"`
	// 商品名称
	SkuName string `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 业务主订单id
	PTradeId string `json:"p_trade_id,omitempty" xml:"p_trade_id,omitempty"`
	// 门店名称
	ShopName string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
	// 门店编码
	ShopCode string `json:"shop_code,omitempty" xml:"shop_code,omitempty"`
	// 交易时间，用户实际下单时间，格式：HH:mm:ss
	TradeTime string `json:"trade_time,omitempty" xml:"trade_time,omitempty"`
	// 业务日期，用户实际下单日期，格式：yyyyMMdd
	BizDate string `json:"biz_date,omitempty" xml:"biz_date,omitempty"`
	// 业务主键
	BizUk string `json:"biz_uk,omitempty" xml:"biz_uk,omitempty"`
	// 未税销售净额
	UntaxSaleTotalAmount int64 `json:"untax_sale_total_amount,omitempty" xml:"untax_sale_total_amount,omitempty"`
	// 含税销售净额
	SaleTotalAmount int64 `json:"sale_total_amount,omitempty" xml:"sale_total_amount,omitempty"`
	// 未税优惠金额
	UntaxDiscountAmount int64 `json:"untax_discount_amount,omitempty" xml:"untax_discount_amount,omitempty"`
	// 含税优惠金额
	DiscountAmount int64 `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	// 未税金额
	UntaxAmount int64 `json:"untax_amount,omitempty" xml:"untax_amount,omitempty"`
	// 含税金额
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 含税商品单价（元）
	UnitPrice int64 `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
	// 交易数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 交易类型编码      * 88 - 销售      * 99 - 退款
	TradeTypeCode int64 `json:"trade_type_code,omitempty" xml:"trade_type_code,omitempty"`
}

var poolFinanceOrderDetail = sync.Pool{
	New: func() any {
		return new(FinanceOrderDetail)
	},
}

// GetFinanceOrderDetail() 从对象池中获取FinanceOrderDetail
func GetFinanceOrderDetail() *FinanceOrderDetail {
	return poolFinanceOrderDetail.Get().(*FinanceOrderDetail)
}

// ReleaseFinanceOrderDetail 释放FinanceOrderDetail
func ReleaseFinanceOrderDetail(v *FinanceOrderDetail) {
	v.Currency = ""
	v.TaxRate = ""
	v.SaleChannel = ""
	v.SaleSource = ""
	v.TradeType = ""
	v.SkuName = ""
	v.SkuCode = ""
	v.PTradeId = ""
	v.ShopName = ""
	v.ShopCode = ""
	v.TradeTime = ""
	v.BizDate = ""
	v.BizUk = ""
	v.UntaxSaleTotalAmount = 0
	v.SaleTotalAmount = 0
	v.UntaxDiscountAmount = 0
	v.DiscountAmount = 0
	v.UntaxAmount = 0
	v.Amount = 0
	v.UnitPrice = 0
	v.Quantity = 0
	v.TradeTypeCode = 0
	poolFinanceOrderDetail.Put(v)
}
