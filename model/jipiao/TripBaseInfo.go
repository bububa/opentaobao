package jipiao

// TripBaseInfo 结构体
type TripBaseInfo struct {
	// 订单创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 订单表最近一次修改时间
	ModifyTime string `json:"modify_time,omitempty" xml:"modify_time,omitempty"`
	// 订单最晚支付时间
	PayLatestTime string `json:"pay_latest_time,omitempty" xml:"pay_latest_time,omitempty"`
	// 佣金，单位：分
	Commission string `json:"commission,omitempty" xml:"commission,omitempty"`
	// 联系人姓名
	RelationName string `json:"relation_name,omitempty" xml:"relation_name,omitempty"`
	// 联系人手机号
	RelationMobile string `json:"relation_mobile,omitempty" xml:"relation_mobile,omitempty"`
	// 联系人备用电话
	RelationPhoneBak string `json:"relation_phone_bak,omitempty" xml:"relation_phone_bak,omitempty"`
	// 联系人邮箱
	RelationEmail string `json:"relation_email,omitempty" xml:"relation_email,omitempty"`
	// 扩展信息
	Extra string `json:"extra,omitempty" xml:"extra,omitempty"`
	// 支付宝交易号
	AlipayTradeNo string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	// 淘宝机票订单id
	OrderId int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 航程类型：0，单程；1，往返；
	TripType int64 `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
	// 订单状态定义==&gt; 0:未付款;(注：支持保险分润时，订单已付款，保险未付款也为0)1:处理中;2:确定出票;3:预定成功;4:预定失败;5:处理中超时失效;6:支付超时失效;7:买家取消
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 订单支付状态==&gt; 0:未付款;1:已付款，冻结买家定金;2:付款给卖家;3:解冻定金给买家;4:已扣佣;5:交易关闭;6:确认可支付;7:确认不可支付
	PayStatus int64 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// 总金额，所有乘机人加机建燃油，单位：分
	TotalPrice int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// 是否保险分润订单，有一个乘机人支持分润即为 True
	InsurePromotion bool `json:"insure_promotion,omitempty" xml:"insure_promotion,omitempty"`
	// 是否强制保险订单，有一张票为强制保险即为true
	ForceInsure bool `json:"force_insure,omitempty" xml:"force_insure,omitempty"`
}
