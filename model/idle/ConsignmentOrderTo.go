package idle

// ConsignmentOrderTo 结构体
type ConsignmentOrderTo struct {
	// 订单创建时间，格式为"yyyy-MM-dd HH:mm:ss"
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 支付方式。1:支付宝现金 2：天猫红包
	IdlePayType string `json:"idle_pay_type,omitempty" xml:"idle_pay_type,omitempty"`
	// 订单环境。online：线上环境 pre：测试环境
	Env string `json:"env,omitempty" xml:"env,omitempty"`
	// 退货原因
	RefundReason string `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
	// 买家关闭原因
	BuyerCloseReason string `json:"buyer_close_reason,omitempty" xml:"buyer_close_reason,omitempty"`
	// 卖家关闭原因
	SellerCloseReason string `json:"seller_close_reason,omitempty" xml:"seller_close_reason,omitempty"`
	// 评价内容
	RateContent string `json:"rate_content,omitempty" xml:"rate_content,omitempty"`
	// 评价等级。1 好评,0 默认中评
	RateGrade string `json:"rate_grade,omitempty" xml:"rate_grade,omitempty"`
	// 预留json格式渠道业务数据，如无则为空。比如,ship=1:不需要发货,weight=5-15代表5-15kg,userLevel=vip代表免议价用户
	ChannelData string `json:"channel_data,omitempty" xml:"channel_data,omitempty"`
	// 子渠道
	SubChannel string `json:"sub_channel,omitempty" xml:"sub_channel,omitempty"`
	// 渠道
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
	// 村
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 区
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 估价ID
	ApprizeId string `json:"apprize_id,omitempty" xml:"apprize_id,omitempty"`
	// 卖家真实姓名
	SellerRealName string `json:"seller_real_name,omitempty" xml:"seller_real_name,omitempty"`
	// 取件时间，格式为"yyyy-MM-dd HH:mm:ss"
	ShipTime string `json:"ship_time,omitempty" xml:"ship_time,omitempty"`
	// 取件类型 1：顺丰 2：上门取件
	ShipType string `json:"ship_type,omitempty" xml:"ship_type,omitempty"`
	// 卖家手机号码
	SellerPhone string `json:"seller_phone,omitempty" xml:"seller_phone,omitempty"`
	// 卖家取件地址/收货地址
	SellerAddress string `json:"seller_address,omitempty" xml:"seller_address,omitempty"`
	// 估价金额，单位分，精确到元
	ApprizeAmount int64 `json:"apprize_amount,omitempty" xml:"apprize_amount,omitempty"`
	// 卖家支付宝账号
	SellerAlipayAccount string `json:"seller_alipay_account,omitempty" xml:"seller_alipay_account,omitempty"`
	// 卖家支付宝用户ID
	SellerAlipayUserId string `json:"seller_alipay_user_id,omitempty" xml:"seller_alipay_user_id,omitempty"`
	// 卖家昵称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 买家昵称
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 订单二级状态,一级状态的子状态,对于没有二级状态的场景该字段为空。一级状态为2已取件: 21:已取件; 22:已收件; 一级状态为3已质检: 31:已质检; 32:用户已确认; 201:一次挂拍; 一级状态为20竞拍中: 202:一次竞拍中; 203:一次竞拍成交; 204:一次拍卖违约; 205:一次竞拍流拍; 211:二次挂拍; 212:二次竞拍中; 213:二次竞拍成交; 214:二次拍卖违约; 215:二次竞拍流拍; 一级状态为5服务商确认交易完成: 51:拍卖成功/订单成功; 58:回收商确认交易/拍卖流拍成交; 59:服务商(兜底)确认交易/支付;
	OrderSubStatus int64 `json:"order_sub_status,omitempty" xml:"order_sub_status,omitempty"`
	// 订单一级状态。1:已下单; 2:已取件; 3:已质检; 20:竞拍中; 5:交易成功; 6:卖家已评价; 7:服务商已评价; 100:申请退货; 101:已退货; 102:卖家取消订单关闭; 103:服务商取消订单关闭;
	OrderStatus int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 闲鱼订单ID
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 回收商appKey
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// json格式服务费计算规则，单位为分，精确到元：最终成交金额*percent%，最高为max金额，最低为min金额。
	ServiceRule string `json:"service_rule,omitempty" xml:"service_rule,omitempty"`
}
