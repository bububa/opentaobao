package idle

// ConsignmentV2OrderTo 结构体
type ConsignmentV2OrderTo struct {
	// 回收商appKey
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 估价金额，单位分，精确到元
	ApprizeAmount string `json:"apprize_amount,omitempty" xml:"apprize_amount,omitempty"`
	// 估价ID
	ApprizeId string `json:"apprize_id,omitempty" xml:"apprize_id,omitempty"`
	// 区
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 闲鱼订单ID
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 买家关闭原因
	BuyerCloseReason string `json:"buyer_close_reason,omitempty" xml:"buyer_close_reason,omitempty"`
	// 渠道
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
	// 预留json格式渠道业务数据，如无则为空。比如,ship=1:不需要发货,weight=5-15代表5-15kg,userLevel=vip代表免议价用户
	ChannelData string `json:"channel_data,omitempty" xml:"channel_data,omitempty"`
	// 市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 用户定价，精确到分
	ConfirmFee string `json:"confirm_fee,omitempty" xml:"confirm_fee,omitempty"`
	// 估价范围，上下限用分号分隔。
	ConsignmentMaxDealAmt string `json:"consignment_max_deal_amt,omitempty" xml:"consignment_max_deal_amt,omitempty"`
	// 服务商最终支付金额，精确到分。成交价减服务费
	ConsignmentPayAmount string `json:"consignment_pay_amount,omitempty" xml:"consignment_pay_amount,omitempty"`
	// 正向服务费，精确到分
	ConsignmentServiceFee string `json:"consignment_service_fee,omitempty" xml:"consignment_service_fee,omitempty"`
	// 村
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 成色，几成新
	Degree string `json:"degree,omitempty" xml:"degree,omitempty"`
	// 下单环境
	Env string `json:"env,omitempty" xml:"env,omitempty"`
	// 订单创建时间，格式为"yyyy-MM-dd HH:mm:ss"
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 订单一级状态
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 订单二级状态
	OrderSubStatus string `json:"order_sub_status,omitempty" xml:"order_sub_status,omitempty"`
	// 省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 卖家定价,分
	QaAmount string `json:"qa_amount,omitempty" xml:"qa_amount,omitempty"`
	// 评价内容
	RateContent string `json:"rate_content,omitempty" xml:"rate_content,omitempty"`
	// 评价等级。1 好评,0 默认中评
	RateGrade string `json:"rate_grade,omitempty" xml:"rate_grade,omitempty"`
	// 回收参考价格，分
	ReferenceRecyclePrice string `json:"reference_recycle_price,omitempty" xml:"reference_recycle_price,omitempty"`
	// 逆向发货物流单
	RefundMailNo string `json:"refund_mail_no,omitempty" xml:"refund_mail_no,omitempty"`
	// 鉴定报告url
	ReportUrl string `json:"report_url,omitempty" xml:"report_url,omitempty"`
	// 逆向服务费，分
	ReverseServiceFee string `json:"reverse_service_fee,omitempty" xml:"reverse_service_fee,omitempty"`
	// 卖家取件地址/收货地址
	SellerAddress string `json:"seller_address,omitempty" xml:"seller_address,omitempty"`
	// 卖家支付宝账号，服务商打款用
	SellerAlipayAccount string `json:"seller_alipay_account,omitempty" xml:"seller_alipay_account,omitempty"`
	// 卖家支付宝账号，服务商打款用
	SellerAlipayUserId string `json:"seller_alipay_user_id,omitempty" xml:"seller_alipay_user_id,omitempty"`
	// 卖家关闭原因
	SellerCloseReason string `json:"seller_close_reason,omitempty" xml:"seller_close_reason,omitempty"`
	// 卖家昵称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 卖家上门取件手机号码
	SellerPhone string `json:"seller_phone,omitempty" xml:"seller_phone,omitempty"`
	// 卖家上门取件名称
	SellerRealName string `json:"seller_real_name,omitempty" xml:"seller_real_name,omitempty"`
	// 逆向卖家退回地址
	SellerReceiptAddress string `json:"seller_receipt_address,omitempty" xml:"seller_receipt_address,omitempty"`
	// 逆向卖家退回取件名称
	SellerReceiptName string `json:"seller_receipt_name,omitempty" xml:"seller_receipt_name,omitempty"`
	// 逆向卖家退回手机号
	SellerReceiptPhone string `json:"seller_receipt_phone,omitempty" xml:"seller_receipt_phone,omitempty"`
	// json格式服务费计算规则，单位为分，精确到元：最终成交金额*percent%，最高为max金额，最低为min金额。
	ServiceRule string `json:"service_rule,omitempty" xml:"service_rule,omitempty"`
	// 上门取件物流单号
	ShipMailNo string `json:"ship_mail_no,omitempty" xml:"ship_mail_no,omitempty"`
	// 上门取件时间
	ShipTime string `json:"ship_time,omitempty" xml:"ship_time,omitempty"`
	// 子渠道
	SubChannel string `json:"sub_channel,omitempty" xml:"sub_channel,omitempty"`
	// 验货结论摘要
	Summary string `json:"summary,omitempty" xml:"summary,omitempty"`
}
