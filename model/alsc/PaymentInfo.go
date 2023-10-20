package alsc

import (
	"sync"
)

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
	// 支付状态：  INITIAL(&#34;INITIAL&#34;, &#34;初始状态&#34;),  WAIT_PAY(&#34;WAIT_PAY&#34;, &#34;待支付状态&#34;),  SUCCESS(&#34;SUCCESS&#34;, &#34;支付成功状态&#34;),  CANCEL(&#34;CANCEL&#34;, &#34;支付冲正状态&#34;),  CLOSE(&#34;CLOSE&#34;, &#34;关闭状态&#34;);
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

var poolPaymentInfo = sync.Pool{
	New: func() any {
		return new(PaymentInfo)
	},
}

// GetPaymentInfo() 从对象池中获取PaymentInfo
func GetPaymentInfo() *PaymentInfo {
	return poolPaymentInfo.Get().(*PaymentInfo)
}

// ReleasePaymentInfo 释放PaymentInfo
func ReleasePaymentInfo(v *PaymentInfo) {
	v.FeeDetailList = v.FeeDetailList[:0]
	v.OutPaymentNo = ""
	v.PayChannel = ""
	v.PayMethod = ""
	v.PaymentStatus = ""
	v.PaymentTime = ""
	v.OrderTotalAmount = 0
	v.OutPayer = nil
	v.SellerReceiveAmount = 0
	v.UserPayAmount = 0
	poolPaymentInfo.Put(v)
}
