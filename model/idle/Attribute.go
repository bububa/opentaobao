package idle

// Attribute 结构体
type Attribute struct {
	// 结构化一级结论枚举值 新品：&#39;CERTIFIED_FLAW&#39;：真货有瑕疵；&#39;CERTIFIED_NO_FLAW&#39;：真货无瑕疵；&#39;FAKE&#39;：假货。 二手：&#39;CERTIFIED_FLAW&#39;：真货有瑕疵；&#39;FAKE&#39;：假货 主状态为4 ac已收货后，服务商调用履约接口传入
	Conclusion1 string `json:"conclusion1,omitempty" xml:"conclusion1,omitempty"`
	// 结构化二级结论枚举值 新品：无需该字段 二手：&#39;SELLER_FAULT&#39;：卖家责任；&#39;SELLER_NO_FAULT&#39;：卖家无责 主状态为4 ac已收货后，服务商调用履约接口传入
	Conclusion2 string `json:"conclusion2,omitempty" xml:"conclusion2,omitempty"`
	// 成色，奢侈品类需要填写&#34; example=&#34;99新&#34;。主状态为4 ac已收货后，服务商调用履约接口传入
	Degree string `json:"degree,omitempty" xml:"degree,omitempty"`
	// 鉴定报告URL地址。主状态为4 ac已收货后，服务商调用履约接口传入
	ReportUrl string `json:"report_url,omitempty" xml:"report_url,omitempty"`
	// 验货结论，真假、有无瑕疵等。主状态为4 ac已收货后，服务商调用履约接口传入
	Summary string `json:"summary,omitempty" xml:"summary,omitempty"`
	// 逆向流程发货给卖家的单号。买家发货后的任意一个逆向状态 服务商调用履约接口传入
	Ac2sellerMailNo string `json:"ac2seller_mail_no,omitempty" xml:"ac2seller_mail_no,omitempty"`
	// 发货给买家的单号。主状态为6 买家确认购买，服务商调用履约接口传入
	Ac2buyerMailNo string `json:"ac2buyer_mail_no,omitempty" xml:"ac2buyer_mail_no,omitempty"`
	// 验货项结论。keyId表示验货项Id，比如1001可能代表内存，valueId指出验货值的id，比如1可能代表存储64G的id，consistent代表实际此值是否为真。关于keyId、valueId的取值由其他对接文档给出
	IdleAppraiseCheckpointsResult string `json:"idle_appraise_checkpoints_result,omitempty" xml:"idle_appraise_checkpoints_result,omitempty"`
	// 拒绝识别场景原因枚举值
	RefuseReasonCode string `json:"refuse_reason_code,omitempty" xml:"refuse_reason_code,omitempty"`
	// 帮卖打款卖家支付宝流水号
	PlatformSalePaidZfbId string `json:"platform_sale_paid_zfb_id,omitempty" xml:"platform_sale_paid_zfb_id,omitempty"`
	// 支付订单号。orderStatus=5时必须上送
	PayOrderId string `json:"pay_order_id,omitempty" xml:"pay_order_id,omitempty"`
	// 拍卖订单违约金额，单位分，精确到元。orderSubStatus=204或者214时必须上送；orderStatus=5时必须上送(违约金总额)
	ForfeitFee string `json:"forfeit_fee,omitempty" xml:"forfeit_fee,omitempty"`
	// 帮卖服务费金额，单位分，精确到元。orderStatus=5时必须上送
	ServiceFee string `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	// 成交金额，单位分，精确到元。orderSubStatus=31时表示相同机况下市场回收价格;orderSubStatus=203或者213,表示拍卖成交金额;orderStatus=5时,表示最终成交金额,必须上送.
	DealAmt string `json:"deal_amt,omitempty" xml:"deal_amt,omitempty"`
	// 节点到期时间，下一节点开始时间（例如订单进入挂拍状态，则表示挂拍结束时间即竞拍开始时间），格式为“yyyy-MM-dd HH:mm:s”。orderSubStatus为如下值时必须上送：201,211表示竞拍开始时间; 202,212表示竞拍结束时间; 203,213表示拍卖买家付款截止时间;
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 拍卖订单号。orderSubStatus=201或者211时上送
	AuctionId string `json:"auction_id,omitempty" xml:"auction_id,omitempty"`
	// 关闭订单原因。orderSubStatus=102时表示卖家取消原因；orderSubStatus=103时表示买家取消原因；
	CloseReason string `json:"close_reason,omitempty" xml:"close_reason,omitempty"`
	// 质检金额（拍卖底价），单位为分，精确到元。orderSubStatus=31时上送
	ConfirmAmt string `json:"confirm_amt,omitempty" xml:"confirm_amt,omitempty"`
	// 快递单号，orderSubStatus=21表示服务商取件快递单号，orderSubStatus =101时表示服务商退货快递单号
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 实际支付金额，单位分，精确到元。orderStatus=5时必须上送
	PayAmt string `json:"pay_amt,omitempty" xml:"pay_amt,omitempty"`
	// 历史最高成交金额，单位分，精确到元。orderSubStatus=31时上送
	MaxDealAmt string `json:"max_deal_amt,omitempty" xml:"max_deal_amt,omitempty"`
	// 评价内容。order_status=7时候填写
	RateContent string `json:"rate_content,omitempty" xml:"rate_content,omitempty"`
	// 评价等级。order_status=7时候填写 1：好评 2：中评
	RateGrade string `json:"rate_grade,omitempty" xml:"rate_grade,omitempty"`
	// 最终成交价(分)
	ConfirmFee string `json:"confirm_fee,omitempty" xml:"confirm_fee,omitempty"`
	// 上门取件订单号
	Seller2AcMailNo string `json:"seller_2_ac_mail_no,omitempty" xml:"seller_2_ac_mail_no,omitempty"`
	// 参考回收价(分)
	ReferenceRecyclePrice string `json:"reference_recycle_price,omitempty" xml:"reference_recycle_price,omitempty"`
	// 逆向退给卖家运单号
	RefundMailNo string `json:"refund_mail_no,omitempty" xml:"refund_mail_no,omitempty"`
	// 服务商发布的商品id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 卖家逆向收货地址id
	SellerReceiptAddressId string `json:"seller_receipt_address_id,omitempty" xml:"seller_receipt_address_id,omitempty"`
	// order_status=103、101时候填写 关闭原因描述
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// order_status=104时填写卖家还欠费金额单位分
	ArrearageFee string `json:"arrearage_fee,omitempty" xml:"arrearage_fee,omitempty"`
	// 字段已废弃，order_status=1时传递支付宝的签约号
	AgreementNo string `json:"agreement_no,omitempty" xml:"agreement_no,omitempty"`
	// 字段已废弃，order_status=1时传递,支付宝用户id
	AlipayUserId string `json:"alipay_user_id,omitempty" xml:"alipay_user_id,omitempty"`
	// 字段已废弃，order_status=104时传递用户还款链接
	ArrearageLink string `json:"arrearage_link,omitempty" xml:"arrearage_link,omitempty"`
	// 业务场景定义的数量，如旧衣回收重量,order_status=3时候传递，具体含义根据不同业务场景决定
	Quantity string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 支付宝交易号，在支付预付款(推送status=8)和推送(status=5)分别提供相应打款的交易号
	AlipayTradeNo string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	// 联系人姓名，上门回收填写,order_status=1
	ContactName string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 联系人电话，上门回收填写,order_status=1
	ContactMobile string `json:"contact_mobile,omitempty" xml:"contact_mobile,omitempty"`
	// 手淘商家的淘宝账号id
	RecycleSupplierId string `json:"recycle_supplier_id,omitempty" xml:"recycle_supplier_id,omitempty"`
	// 商家收到货品时间
	IsvReceiveTime string `json:"isv_receive_time,omitempty" xml:"isv_receive_time,omitempty"`
	// 质检型号名称（若存在不一致情况，需要了解实际检测的质检型号，若实际型号在已挂载的spuid中则传spuid，否则传&#34;其他&#34;）
	RealModel string `json:"real_model,omitempty" xml:"real_model,omitempty"`
	// 提交型号和质检型号是否一致
	SpuMatch string `json:"spu_match,omitempty" xml:"spu_match,omitempty"`
	// order_status=103、101时候填写 关闭原因code。QA_STAFF_NOT_VISIT_HOME(&#34;质检员未上门取件&#34;), SELLER_CAN_NOT_CONTACT(&#34;用户无法联系&#34;), SELLER_NOT_COME_STORE_AT_TIME(&#34;用户未按时到店&#34;), CANCEL_BY_SELLER_DEMAND(&#34;用户要求不回收了&#34;), QA_NOT_QUALIFIED(&#34;不符合服务商质检要求&#34;), OTHER(&#34;其他原因&#34;)
	CloseReasonCode string `json:"close_reason_code,omitempty" xml:"close_reason_code,omitempty"`
	// 逆向退回邮费，单位分
	RefundLogisticsFee string `json:"refund_logistics_fee,omitempty" xml:"refund_logistics_fee,omitempty"`
	// 补偿给卖家的费用
	CompensationFee string `json:"compensation_fee,omitempty" xml:"compensation_fee,omitempty"`
	// 是否允许用户修改地址
	AgreeUseAddressChange string `json:"agree_use_address_change,omitempty" xml:"agree_use_address_change,omitempty"`
}
