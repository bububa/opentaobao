package bus

import (
	"sync"
)

// RefundAccountInDetail 结构体
type RefundAccountInDetail struct {
	// 收款账户支付宝ID 必传
	AlipayAccount string `json:"alipay_account,omitempty" xml:"alipay_account,omitempty"`
	// 收款装好支付宝账号，注意和上面的支付宝ID 要对应好
	AlipayAccountId string `json:"alipay_account_id,omitempty" xml:"alipay_account_id,omitempty"`
	// 退款批次号须唯一
	RefundBatchNo string `json:"refund_batch_no,omitempty" xml:"refund_batch_no,omitempty"`
	// 分为单位;退多少钱
	RefundAmount int64 `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
}

var poolRefundAccountInDetail = sync.Pool{
	New: func() any {
		return new(RefundAccountInDetail)
	},
}

// GetRefundAccountInDetail() 从对象池中获取RefundAccountInDetail
func GetRefundAccountInDetail() *RefundAccountInDetail {
	return poolRefundAccountInDetail.Get().(*RefundAccountInDetail)
}

// ReleaseRefundAccountInDetail 释放RefundAccountInDetail
func ReleaseRefundAccountInDetail(v *RefundAccountInDetail) {
	v.AlipayAccount = ""
	v.AlipayAccountId = ""
	v.RefundBatchNo = ""
	v.RefundAmount = 0
	poolRefundAccountInDetail.Put(v)
}
