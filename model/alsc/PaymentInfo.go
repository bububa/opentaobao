package alsc

// PaymentInfo 结构体
type PaymentInfo struct {
	// 费用明细
	FeeDetailList []FeeDetail `json:"fee_detail_list,omitempty" xml:"fee_detail_list>fee_detail,omitempty"`
	// 外部付款单号
	OutPaymentNo string `json:"out_payment_no,omitempty" xml:"out_payment_no,omitempty"`
	// 支付宝-ALIPAY  微信-WECHAT  其他-OTHER
	PayChannel string `json:"pay_channel,omitempty" xml:"pay_channel,omitempty"`
	// 现金-CASH,  代扣-DEDUCT,  银行卡支付-BANK,  储值-STOREDVALUE,  微信-WECHAT,  支付宝 -ALIPAY,  积分-POINTS,  线下券-VOUCHER,  其他方式-OTHER;
	PayMethod string `json:"pay_method,omitempty" xml:"pay_method,omitempty"`
	// 支付状态：  INITIAL("INITIAL", "初始状态"),  WAIT_PAY("WAIT_PAY", "待支付状态"),  SUCCESS("SUCCESS", "支付成功状态"),  CANCEL("CANCEL", "支付冲正状态"),  CLOSE("CLOSE", "关闭状态");
	PaymentStatus string `json:"payment_status,omitempty" xml:"payment_status,omitempty"`
	// 支付时间
	PaymentTime string `json:"payment_time,omitempty" xml:"payment_time,omitempty"`
	// 订单总金额
	OrderTotalAmount int64 `json:"order_total_amount,omitempty" xml:"order_total_amount,omitempty"`
	// 外部付款方基础信息
	OutPayer *OrderUser `json:"out_payer,omitempty" xml:"out_payer,omitempty"`
	// 商家实收
	SellerReceiveAmount int64 `json:"seller_receive_amount,omitempty" xml:"seller_receive_amount,omitempty"`
	// 付款方实付金额
	UserPayAmount int64 `json:"user_pay_amount,omitempty" xml:"user_pay_amount,omitempty"`
}
