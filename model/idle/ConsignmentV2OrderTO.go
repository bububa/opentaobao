package idle

// ConsignmentV2orderTo 结构体
type ConsignmentV2orderTo struct {
	// appkey
	Appkey string `json:"appkey,omitempty" xml:"appkey,omitempty"`
	// 订单id
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 订单主状态
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 订单子状态
	OrderSubStatus string `json:"order_sub_status,omitempty" xml:"order_sub_status,omitempty"`
	// 卖家id
	SellerId string `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 卖家昵称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 卖家支付宝昵称
	SellerAlipayUserId string `json:"seller_alipay_user_id,omitempty" xml:"seller_alipay_user_id,omitempty"`
	// 卖家支付宝账户
	SellerAlipayAccount string `json:"seller_alipay_account,omitempty" xml:"seller_alipay_account,omitempty"`
	// 估价金额
	ApprizeAmount string `json:"apprize_amount,omitempty" xml:"apprize_amount,omitempty"`
	// 卖家地址
	SellerAddress string `json:"seller_address,omitempty" xml:"seller_address,omitempty"`
	// 卖家手机
	SellerPhone string `json:"seller_phone,omitempty" xml:"seller_phone,omitempty"`
	// 上门时间
	ShipTime string `json:"ship_time,omitempty" xml:"ship_time,omitempty"`
	// 卖家真实姓名
	SellerRealName string `json:"seller_real_name,omitempty" xml:"seller_real_name,omitempty"`
	// 估价id
	ApprizeId string `json:"apprize_id,omitempty" xml:"apprize_id,omitempty"`
	// 市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 区
	Area string `json:"area,omitempty" xml:"area,omitempty"`
	// 村
	Country string `json:"country,omitempty" xml:"country,omitempty"`
	// 订单渠道
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
	// 子渠道
	SubChannel string `json:"sub_channel,omitempty" xml:"sub_channel,omitempty"`
	// 渠道下存储的数据
	ChannelData string `json:"channel_data,omitempty" xml:"channel_data,omitempty"`
	// 评价等级 1 好评,0 默认中评
	RateGrade string `json:"rate_grade,omitempty" xml:"rate_grade,omitempty"`
	// 评价内容
	RateContent string `json:"rate_content,omitempty" xml:"rate_content,omitempty"`
	// 卖家取消订单原因
	SellerCloseReason string `json:"seller_close_reason,omitempty" xml:"seller_close_reason,omitempty"`
	// 买家关闭订单原因
	BuyerCloseReason string `json:"buyer_close_reason,omitempty" xml:"buyer_close_reason,omitempty"`
	// 订单产生环境
	Env string `json:"env,omitempty" xml:"env,omitempty"`
	// 订单创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 服务费规则
	ServiceRule string `json:"service_rule,omitempty" xml:"service_rule,omitempty"`
	// 上门取件物流单号
	ShipMailNo string `json:"ship_mail_no,omitempty" xml:"ship_mail_no,omitempty"`
	// 质检结果简述
	Summary string `json:"summary,omitempty" xml:"summary,omitempty"`
	// 质检报告url
	ReportUrl string `json:"report_url,omitempty" xml:"report_url,omitempty"`
	// 成色
	Degree string `json:"degree,omitempty" xml:"degree,omitempty"`
	// 服务费
	ConsignmentServiceFee string `json:"consignment_service_fee,omitempty" xml:"consignment_service_fee,omitempty"`
	// 服务商实际支付金额
	ConsignmentPayAmount string `json:"consignment_pay_amount,omitempty" xml:"consignment_pay_amount,omitempty"`
	// 质检价格
	QaAmount string `json:"qa_amount,omitempty" xml:"qa_amount,omitempty"`
	// 用户出价
	ConsignmentMaxDealAmt string `json:"consignment_max_deal_amt,omitempty" xml:"consignment_max_deal_amt,omitempty"`
	// 参考回收价
	ReferenceRecyclePrice string `json:"reference_recycle_price,omitempty" xml:"reference_recycle_price,omitempty"`
	// 逆向退回地址id
	SellerReceiptAddressId string `json:"seller_receipt_address_id,omitempty" xml:"seller_receipt_address_id,omitempty"`
	// 逆向退回地址 详细信息
	SellerReceiptAddress string `json:"seller_receipt_address,omitempty" xml:"seller_receipt_address,omitempty"`
	// 退回运单号
	RefundMailNo string `json:"refund_mail_no,omitempty" xml:"refund_mail_no,omitempty"`
	// 逆向退回卖家手机号
	SellerReceiptPhone string `json:"seller_receipt_phone,omitempty" xml:"seller_receipt_phone,omitempty"`
	// 逆向退回卖家姓名
	SellerReceiptName string `json:"seller_receipt_name,omitempty" xml:"seller_receipt_name,omitempty"`
	// 逆向处理费
	ReverseServiceFee string `json:"reverse_service_fee,omitempty" xml:"reverse_service_fee,omitempty"`
	// spuId
	SpuId string `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
	// 订单发货类型，1:邮寄，2:上门，3:到店
	ShipType string `json:"ship_type,omitempty" xml:"ship_type,omitempty"`
	// 门店地址id
	StationId string `json:"station_id,omitempty" xml:"station_id,omitempty"`
	// 回收商appKey
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 用户定价，精确到分
	ConfirmFee string `json:"confirm_fee,omitempty" xml:"confirm_fee,omitempty"`
}
