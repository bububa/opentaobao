package aliospay

import (
	"sync"
)

// GetTokenRequest 结构体
type GetTokenRequest struct {
	// 请求唯一id，不可重复，服务端会根据此参数防重放
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// ORDER订单token，OTHER其他token
	TokenType string `json:"token_type,omitempty" xml:"token_type,omitempty"`
	// 订单标题
	Subject string `json:"subject,omitempty" xml:"subject,omitempty"`
	// 业务订单号
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 请求时间戳
	Time string `json:"time,omitempty" xml:"time,omitempty"`
	// 语言,en表示英文，zh表示中文
	Lang string `json:"lang,omitempty" xml:"lang,omitempty"`
	// 周期扣款规则参数，周期扣款订单必传。支付系统会按照这里传入的参数提示用户，并对发起扣款的时间、金额、次数等做相应限制
	PeriodRuleParams string `json:"period_rule_params,omitempty" xml:"period_rule_params,omitempty"`
	// cp服务端支持的协议，目前只支持HTTPS
	ServiceProtocol string `json:"service_protocol,omitempty" xml:"service_protocol,omitempty"`
	// 周期扣款签约结果回调地址
	PeriodSignNotifyUrl string `json:"period_sign_notify_url,omitempty" xml:"period_sign_notify_url,omitempty"`
	// 周期扣款解约结果回调地址
	PeriodUnsignNotifyUrl string `json:"period_unsign_notify_url,omitempty" xml:"period_unsign_notify_url,omitempty"`
	// 支付结果回调通知URL
	PayNotifyUrl string `json:"pay_notify_url,omitempty" xml:"pay_notify_url,omitempty"`
	// CP的商品原始金额，不参与任何计算，仅用于展示原始金额
	OriginalAmount int64 `json:"original_amount,omitempty" xml:"original_amount,omitempty"`
	// 参与优惠计算的金额，用此字段用于让订单中部分金额不参与优惠的计算
	DiscountableAmount int64 `json:"discountable_amount,omitempty" xml:"discountable_amount,omitempty"`
	// 订单总金额
	TotalAmount int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
}

var poolGetTokenRequest = sync.Pool{
	New: func() any {
		return new(GetTokenRequest)
	},
}

// GetGetTokenRequest() 从对象池中获取GetTokenRequest
func GetGetTokenRequest() *GetTokenRequest {
	return poolGetTokenRequest.Get().(*GetTokenRequest)
}

// ReleaseGetTokenRequest 释放GetTokenRequest
func ReleaseGetTokenRequest(v *GetTokenRequest) {
	v.TraceId = ""
	v.TokenType = ""
	v.Subject = ""
	v.BizOrderId = ""
	v.Time = ""
	v.Lang = ""
	v.PeriodRuleParams = ""
	v.ServiceProtocol = ""
	v.PeriodSignNotifyUrl = ""
	v.PeriodUnsignNotifyUrl = ""
	v.PayNotifyUrl = ""
	v.OriginalAmount = 0
	v.DiscountableAmount = 0
	v.TotalAmount = 0
	poolGetTokenRequest.Put(v)
}
