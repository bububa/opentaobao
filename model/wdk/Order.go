package wdk

// Order 结构体
type Order struct {
	// 支付渠道列表
	PayChannels []PayChannel `json:"pay_channels,omitempty" xml:"pay_channels>pay_channel,omitempty"`
	// 子订单列表
	SubOrders []SubOrder `json:"sub_orders,omitempty" xml:"sub_orders>sub_order,omitempty"`
	// 成交时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 门店编码
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 会员卡号
	MemberCardNum string `json:"member_card_num,omitempty" xml:"member_card_num,omitempty"`
	// 收银员编号
	OperatorId string `json:"operator_id,omitempty" xml:"operator_id,omitempty"`
	// 收银员名字
	OperatorName string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
	// merchantCode
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// orderStatus
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// memberPoint
	MemberPoint string `json:"member_point,omitempty" xml:"member_point,omitempty"`
	// 打包时间
	PackageTime string `json:"package_time,omitempty" xml:"package_time,omitempty"`
	// 赠券，格式为 券ID_券金额， 金额单位为分
	GiftCoupon string `json:"gift_coupon,omitempty" xml:"gift_coupon,omitempty"`
	// 一串唯一的字符串
	DutyCode string `json:"duty_code,omitempty" xml:"duty_code,omitempty"`
	// map格式的json字符串，部分key含义如下。wdkMemberValue:手机号或者卡号 wdkMemberSource:carNum or phoneNum memDegrade：1表示在线降级，2表示离线降级
	TradeAttributes string `json:"trade_attributes,omitempty" xml:"trade_attributes,omitempty"`
	// 买家open_uid
	OpenUid string `json:"open_uid,omitempty" xml:"open_uid,omitempty"`
	// 渠道店id
	ShopId string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// 外部门店id
	OutShopId string `json:"out_shop_id,omitempty" xml:"out_shop_id,omitempty"`
	// 外部订单号
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// app或pos
	OrderClient string `json:"order_client,omitempty" xml:"order_client,omitempty"`
	// 业务订单号
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 商品优惠总额
	DiscountAmt int64 `json:"discount_amt,omitempty" xml:"discount_amt,omitempty"`
	// 商品总金额（优惠前）
	OriginalAmt int64 `json:"original_amt,omitempty" xml:"original_amt,omitempty"`
	// 配送费
	PostFee int64 `json:"post_fee,omitempty" xml:"post_fee,omitempty"`
	// 销售类型（正向销售1：逆向销售2。本接口应返回2）
	TrdType int64 `json:"trd_type,omitempty" xml:"trd_type,omitempty"`
	// 会员价优惠总金额
	MemberDiscountAmt int64 `json:"member_discount_amt,omitempty" xml:"member_discount_amt,omitempty"`
	// 淘宝主订单号
	TbBizOrderId int64 `json:"tb_biz_order_id,omitempty" xml:"tb_biz_order_id,omitempty"`
	// 渠道来源
	OrderFrom int64 `json:"order_from,omitempty" xml:"order_from,omitempty"`
	// 支付金额，单位分
	PayAmt int64 `json:"pay_amt,omitempty" xml:"pay_amt,omitempty"`
}
