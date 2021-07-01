package trade

// OrderDiscountBillBo 结构体
type OrderDiscountBillBo struct {
	// 活动ID
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 活动名称
	ActivityName string `json:"activity_name,omitempty" xml:"activity_name,omitempty"`
	// 活动类型 1：活动 2：券
	ActivityType int64 `json:"activity_type,omitempty" xml:"activity_type,omitempty"`
	// 购买数量
	BuyQuantity int64 `json:"buy_quantity,omitempty" xml:"buy_quantity,omitempty"`
	// 优惠金额
	DiscountFee int64 `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
	// 优惠件数
	DiscountQuantity int64 `json:"discount_quantity,omitempty" xml:"discount_quantity,omitempty"`
	// 补差类型
	DiscountType int64 `json:"discount_type,omitempty" xml:"discount_type,omitempty"`
	// 主键id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 主订单号
	MainOrderId string `json:"main_order_id,omitempty" xml:"main_order_id,omitempty"`
	// 商家编码
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 商家优惠补差金额
	MerchantDiscountFee int64 `json:"merchant_discount_fee,omitempty" xml:"merchant_discount_fee,omitempty"`
	// 交易状态
	OrderStatus int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 外部订单号
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 商品编码
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 商品名称
	SkuName string `json:"sku_name,omitempty" xml:"sku_name,omitempty"`
	// 经营店id
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 子单号
	SubOrderId string `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// 淘鲜达优惠金额
	TxdDiscountFee int64 `json:"txd_discount_fee,omitempty" xml:"txd_discount_fee,omitempty"`
	// 业务时间
	BizTime string `json:"biz_time,omitempty" xml:"biz_time,omitempty"`
	// 订单渠道
	OrderChannel int64 `json:"order_channel,omitempty" xml:"order_channel,omitempty"`
}
