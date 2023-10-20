package wdk

// OrderSuccessRequest 结构体
type OrderSuccessRequest struct {
	// 子单列表
	SubInfoList []OrderDeliveryBo `json:"sub_info_list,omitempty" xml:"sub_info_list>order_delivery_bo,omitempty"`
	// 外部订单号
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 订单来源
	OrderSource string `json:"order_source,omitempty" xml:"order_source,omitempty"`
	// 下单终端
	OrderTerminal string `json:"order_terminal,omitempty" xml:"order_terminal,omitempty"`
	// 一级渠道
	FirstChannel string `json:"first_channel,omitempty" xml:"first_channel,omitempty"`
	// 二级渠道
	SecondChannel string `json:"second_channel,omitempty" xml:"second_channel,omitempty"`
	// 渠道店
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 经营店id
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 创单时间
	OrderCreateTime string `json:"order_create_time,omitempty" xml:"order_create_time,omitempty"`
	// 支付时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 期望送达时间
	ExpectArriveTime string `json:"expect_arrive_time,omitempty" xml:"expect_arrive_time,omitempty"`
	// 订单支付信息
	PayInfos string `json:"pay_infos,omitempty" xml:"pay_infos,omitempty"`
	// 卖家昵称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 是否是主单 1-是;0-不是
	IsMain int64 `json:"is_main,omitempty" xml:"is_main,omitempty"`
	// 是否是子单 1-是;0-不是
	IsDetail int64 `json:"is_detail,omitempty" xml:"is_detail,omitempty"`
	// 业务类型 2-表示生鲜门店;3-表示B2C
	BusinessType int64 `json:"business_type,omitempty" xml:"business_type,omitempty"`
	// 子业务类型
	SubBusinessType int64 `json:"sub_business_type,omitempty" xml:"sub_business_type,omitempty"`
	// 订单渠道 2-表示线上;3-表示线下
	OrderChannel int64 `json:"order_channel,omitempty" xml:"order_channel,omitempty"`
	// 配送类型 1-表示预约配送;2-表示现场购买
	DeliverType int64 `json:"deliver_type,omitempty" xml:"deliver_type,omitempty"`
	// 送达类型 1-表示及时达;2-表示定时达;3-表示极速达;4-表示无需配送
	ArriveType int64 `json:"arrive_type,omitempty" xml:"arrive_type,omitempty"`
	// 订单状态
	OrderStatus int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 订单原价(分为单位)
	OriginFee int64 `json:"origin_fee,omitempty" xml:"origin_fee,omitempty"`
	// 支付金额(分为单位)
	PayFee int64 `json:"pay_fee,omitempty" xml:"pay_fee,omitempty"`
	// 优惠金额(分为单位)
	DiscountFee int64 `json:"discount_fee,omitempty" xml:"discount_fee,omitempty"`
	// 运费(分为单位)
	PostFee int64 `json:"post_fee,omitempty" xml:"post_fee,omitempty"`
	// 打包费(分为单位)
	PackageFee int64 `json:"package_fee,omitempty" xml:"package_fee,omitempty"`
	// 平台折扣费(分为单位)
	PlatformDiscountFee int64 `json:"platform_discount_fee,omitempty" xml:"platform_discount_fee,omitempty"`
	// 商家折扣费(分为单位)
	MerchantDiscountFee int64 `json:"merchant_discount_fee,omitempty" xml:"merchant_discount_fee,omitempty"`
	// 买家信息
	BuyerInfo *OrderPayInfoBo `json:"buyer_info,omitempty" xml:"buyer_info,omitempty"`
	// 订单配送信息
	DeliveryInfo *OrderBuyerInfoBo `json:"delivery_info,omitempty" xml:"delivery_info,omitempty"`
	// 卖家id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 订单来源
	OrderFrom int64 `json:"order_from,omitempty" xml:"order_from,omitempty"`
}
