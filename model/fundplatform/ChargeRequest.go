package fundplatform

import (
	"sync"
)

// ChargeRequest 结构体
type ChargeRequest struct {
	// 描述信息
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 幂等字段,最多32位,必须保证同一请求幂等号唯一。不同请求幂等号不同
	OutBizId string `json:"out_biz_id,omitempty" xml:"out_biz_id,omitempty"`
	// 付款支付宝账户email
	PayerAlipayEmail string `json:"payer_alipay_email,omitempty" xml:"payer_alipay_email,omitempty"`
	// 付款支付宝账户编码
	PayerAlipayNo string `json:"payer_alipay_no,omitempty" xml:"payer_alipay_no,omitempty"`
	// 充值金额，单位分
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 业务类型,由资金平台分配,一般为9位数字
	SubBizType int64 `json:"sub_biz_type,omitempty" xml:"sub_biz_type,omitempty"`
	// 用户ID,两个userId请保持一致
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

var poolChargeRequest = sync.Pool{
	New: func() any {
		return new(ChargeRequest)
	},
}

// GetChargeRequest() 从对象池中获取ChargeRequest
func GetChargeRequest() *ChargeRequest {
	return poolChargeRequest.Get().(*ChargeRequest)
}

// ReleaseChargeRequest 释放ChargeRequest
func ReleaseChargeRequest(v *ChargeRequest) {
	v.Description = ""
	v.OutBizId = ""
	v.PayerAlipayEmail = ""
	v.PayerAlipayNo = ""
	v.Amount = 0
	v.SubBizType = 0
	v.UserId = 0
	poolChargeRequest.Put(v)
}
