package wdk

// Suborders 结构体
type Suborders struct {
	// 外部子订单id
	OutSubOrderId string `json:"out_sub_order_id,omitempty" xml:"out_sub_order_id,omitempty"`
	// 商品code
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 商品名称
	AuctionTitle string `json:"auction_title,omitempty" xml:"auction_title,omitempty"`
	// 库存单位购买数量
	BuyAmountStock string `json:"buy_amount_stock,omitempty" xml:"buy_amount_stock,omitempty"`
	// 销售单位
	SaleUnit string `json:"sale_unit,omitempty" xml:"sale_unit,omitempty"`
	// 库存单位
	StockUnit string `json:"stock_unit,omitempty" xml:"stock_unit,omitempty"`
	// 五道口子订单id
	BizSubOrderId int64 `json:"biz_sub_order_id,omitempty" xml:"biz_sub_order_id,omitempty"`
	// 商品价格
	AuctionPrice int64 `json:"auction_price,omitempty" xml:"auction_price,omitempty"`
	// 销售单位购买数量
	BuyAmountSale int64 `json:"buy_amount_sale,omitempty" xml:"buy_amount_sale,omitempty"`
	// 子订单用户支付金额
	PayFee int64 `json:"pay_fee,omitempty" xml:"pay_fee,omitempty"`
	// 子订单原价
	OriginFee int64 `json:"origin_fee,omitempty" xml:"origin_fee,omitempty"`
	// 子订单优惠金额
	DiscountFee int64 `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
	// 商家优惠分摊
	DiscountMerchantFee int64 `json:"discount_merchant_fee,omitempty" xml:"discount_merchant_fee,omitempty"`
	// 平台优惠分摊
	DiscountPlatformFee int64 `json:"discount_platform_fee,omitempty" xml:"discount_platform_fee,omitempty"`
	// 子单价格
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
}
