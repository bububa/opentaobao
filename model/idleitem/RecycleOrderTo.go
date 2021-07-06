package idleitem

// RecycleOrderTo 结构体
type RecycleOrderTo struct {
	// 回收商appkey
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 估价Id
	ApprizeId string `json:"apprize_id,omitempty" xml:"apprize_id,omitempty"`
	// 区
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 交易订单号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 回收商取消订单原因
	BuyerCloseReason string `json:"buyer_close_reason,omitempty" xml:"buyer_close_reason,omitempty"`
	// 用户取消订单原因
	SellerCloseReason string `json:"seller_close_reason,omitempty" xml:"seller_close_reason,omitempty"`
	// 回收商买家账号
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 渠道信息
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
	// 渠道内的业务数据json格式 比如 ship=1 服装类的不需要发货,weight=5-15 代表服装5-15kg, userLevel=vip 代表免议价用户,sellerRealPhone 淘宝账号绑定的手机号
	ChannelData string `json:"channel_data,omitempty" xml:"channel_data,omitempty"`
	// 城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 村
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 加价券金额,单位分
	CouponFee string `json:"coupon_fee,omitempty" xml:"coupon_fee,omitempty"`
	// 加价券Id（预留）
	CouponId string `json:"coupon_id,omitempty" xml:"coupon_id,omitempty"`
	// 加价券规则（预留）
	CouponRule string `json:"coupon_rule,omitempty" xml:"coupon_rule,omitempty"`
	// onlien:线上环境 pre：测试环境
	Env string `json:"env,omitempty" xml:"env,omitempty"`
	// 订单创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 1:现金支付
	IdlePayType string `json:"idle_pay_type,omitempty" xml:"idle_pay_type,omitempty"`
	// 省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 卖家评价内容
	RateContent string `json:"rate_content,omitempty" xml:"rate_content,omitempty"`
	// 卖家评价等级
	RateGrade string `json:"rate_grade,omitempty" xml:"rate_grade,omitempty"`
	// 卖家申请退回原因
	RefundReason string `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
	// 卖家收货地址
	SellerAddress string `json:"seller_address,omitempty" xml:"seller_address,omitempty"`
	// 卖家淘宝nick
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 卖家联系号码
	SellerPhone string `json:"seller_phone,omitempty" xml:"seller_phone,omitempty"`
	// 卖家真实姓名
	SellerRealName string `json:"seller_real_name,omitempty" xml:"seller_real_name,omitempty"`
	// 上门取件时间
	ShipTime string `json:"ship_time,omitempty" xml:"ship_time,omitempty"`
	// 取件类型
	ShipType string `json:"ship_type,omitempty" xml:"ship_type,omitempty"`
	// 子渠道
	SubChannel string `json:"sub_channel,omitempty" xml:"sub_channel,omitempty"`
	// 支付状态 1：未付款  2：已付款
	PayStatus string `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// 下单后付款时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 订单完结时间
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 追缴单流水Id
	RecoveryPayId string `json:"recovery_pay_id,omitempty" xml:"recovery_pay_id,omitempty"`
	// 追缴成功时间
	RecoveryPayTime string `json:"recovery_pay_time,omitempty" xml:"recovery_pay_time,omitempty"`
	// 少补单流水Id
	MakeupPayId string `json:"makeup_pay_id,omitempty" xml:"makeup_pay_id,omitempty"`
	// 少补成功时间
	MakeupPayTime string `json:"makeup_pay_time,omitempty" xml:"makeup_pay_time,omitempty"`
	// 主订单支付流水Id
	MainPayId string `json:"main_pay_id,omitempty" xml:"main_pay_id,omitempty"`
	// 估价金额,单位分
	ApprizeAmount int64 `json:"apprize_amount,omitempty" xml:"apprize_amount,omitempty"`
	// 信用预付金额,单位分
	CreditPayAmount int64 `json:"credit_pay_amount,omitempty" xml:"credit_pay_amount,omitempty"`
	// 回收订单状态：1:订单创建 2:已上门取件 3:已质检 4:卖家确认交易完成 6:卖家订单已评价 7:回收商订单已评价 100:卖家申请退回 101:货物已退回 102:卖家关闭订单
	OrderStatus int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 质检金额,单位分
	QaAmount int64 `json:"qa_amount,omitempty" xml:"qa_amount,omitempty"`
	// true：是信用预付订单，false：普通订单
	CreditPay bool `json:"credit_pay,omitempty" xml:"credit_pay,omitempty"`
}
