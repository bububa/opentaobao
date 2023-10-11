package wdk

// SubOrderResponse 结构体
type SubOrderResponse struct {
	// 子单活动列表
	Activitys []OrderActivity `json:"activitys,omitempty" xml:"activitys>order_activity,omitempty"`
	// 资金优惠
	FundsDiscounts []OrderFundsDiscount `json:"funds_discounts,omitempty" xml:"funds_discounts>order_funds_discount,omitempty"`
	// 渠道子订单编码
	OutSubOrderId string `json:"out_sub_order_id,omitempty" xml:"out_sub_order_id,omitempty"`
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 商品名称
	SkuName string `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	// 销售单位
	SaleUnit string `json:"sale_unit,omitempty" xml:"sale_unit,omitempty"`
	// 库存单位
	StockUnit string `json:"stock_unit,omitempty" xml:"stock_unit,omitempty"`
	// 商品条码
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// 是否赠品 0=普通品/1=赠品
	GiftFlag string `json:"gift_flag,omitempty" xml:"gift_flag,omitempty"`
	// 是否称重品 0=标品 / 1=称重品
	WeightFlag string `json:"weight_flag,omitempty" xml:"weight_flag,omitempty"`
	// 子订单状态 值枚举同主单状态
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 子订单编码
	BizSubOrderId int64 `json:"biz_sub_order_id,omitempty" xml:"biz_sub_order_id,omitempty"`
	// 商品单价, 分
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 库存单位购买数量
	BuyStockQuantity *BigDecimal `json:"buy_stock_quantity,omitempty" xml:"buy_stock_quantity,omitempty"`
	// 销售单位购买数量
	BuySaleQuantity *BigDecimal `json:"buy_sale_quantity,omitempty" xml:"buy_sale_quantity,omitempty"`
	// 子订单原价
	OriginalFee int64 `json:"original_fee,omitempty" xml:"original_fee,omitempty"`
	// 子订单优惠金额
	DiscountFee int64 `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
	// 商品总重量, g
	Weight int64 `json:"weight,omitempty" xml:"weight,omitempty"`
	// 商家优惠分摊
	DiscountMerchantFee int64 `json:"discount_merchant_fee,omitempty" xml:"discount_merchant_fee,omitempty"`
	// 平台优惠分摊
	DiscountPlatformFee int64 `json:"discount_platform_fee,omitempty" xml:"discount_platform_fee,omitempty"`
}
