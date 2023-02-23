package vaccin

// TradeSubscribeDetailExecuteTopRequest 结构体
type TradeSubscribeDetailExecuteTopRequest struct {
	// 外部商家预约单主键
	IsvSubscribeId string `json:"isv_subscribe_id,omitempty" xml:"isv_subscribe_id,omitempty"`
	// 商家给用户安排预约时间
	SubscribeTime string `json:"subscribe_time,omitempty" xml:"subscribe_time,omitempty"`
	// 商家ID
	MerchantId string `json:"merchant_id,omitempty" xml:"merchant_id,omitempty"`
	// 门店名称
	ShopName string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
	// 取消原因
	CancelReason string `json:"cancel_reason,omitempty" xml:"cancel_reason,omitempty"`
	// 统一社会信用编码
	SocialCreditCode string `json:"social_credit_code,omitempty" xml:"social_credit_code,omitempty"`
	// 商家预约接种开始时间
	SubscribeStartTime string `json:"subscribe_start_time,omitempty" xml:"subscribe_start_time,omitempty"`
	// 商家预约接种结束时间
	SubscribeEndTime string `json:"subscribe_end_time,omitempty" xml:"subscribe_end_time,omitempty"`
	// 针次(1,2,3,4)
	InjectionNum int64 `json:"injection_num,omitempty" xml:"injection_num,omitempty"`
	// 业务订单号
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 预约单状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 接种状态
	InnoculateStatus int64 `json:"innoculate_status,omitempty" xml:"innoculate_status,omitempty"`
}
