package ascpchannel

import (
	"sync"
)

// ExternalCreateRefundOrderRequest 结构体
type ExternalCreateRefundOrderRequest struct {
	// 退款举证图片列表
	ProofPicUrls []string `json:"proof_pic_urls,omitempty" xml:"proof_pic_urls>string,omitempty"`
	// 币种
	CurrencyType string `json:"currency_type,omitempty" xml:"currency_type,omitempty"`
	// 销售订单号
	SaleOrderNo string `json:"sale_order_no,omitempty" xml:"sale_order_no,omitempty"`
	// 外部退款单号
	OutRefundNo string `json:"out_refund_no,omitempty" xml:"out_refund_no,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 外部订单号
	OutOrderNo string `json:"out_order_no,omitempty" xml:"out_order_no,omitempty"`
	// 退款原因
	RefundReason string `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
}

var poolExternalCreateRefundOrderRequest = sync.Pool{
	New: func() any {
		return new(ExternalCreateRefundOrderRequest)
	},
}

// GetExternalCreateRefundOrderRequest() 从对象池中获取ExternalCreateRefundOrderRequest
func GetExternalCreateRefundOrderRequest() *ExternalCreateRefundOrderRequest {
	return poolExternalCreateRefundOrderRequest.Get().(*ExternalCreateRefundOrderRequest)
}

// ReleaseExternalCreateRefundOrderRequest 释放ExternalCreateRefundOrderRequest
func ReleaseExternalCreateRefundOrderRequest(v *ExternalCreateRefundOrderRequest) {
	v.ProofPicUrls = v.ProofPicUrls[:0]
	v.CurrencyType = ""
	v.SaleOrderNo = ""
	v.OutRefundNo = ""
	v.Remark = ""
	v.OutOrderNo = ""
	v.RefundReason = ""
	poolExternalCreateRefundOrderRequest.Put(v)
}
