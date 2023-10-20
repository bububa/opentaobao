package elife

import (
	"sync"
)

// RefundRequest 结构体
type RefundRequest struct {
	// 操作流水号, 商家全系统唯一
	OpId string `json:"op_id,omitempty" xml:"op_id,omitempty"`
	// 收银员名称
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 原正向核销交易的流水号
	OriginalOpId string `json:"original_op_id,omitempty" xml:"original_op_id,omitempty"`
	// 门店编号
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 原正向核销交易的核销码
	PayCode string `json:"pay_code,omitempty" xml:"pay_code,omitempty"`
	// 原正向核销交易的小票信息
	SaleTicket string `json:"sale_ticket,omitempty" xml:"sale_ticket,omitempty"`
	// 金额分
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
}

var poolRefundRequest = sync.Pool{
	New: func() any {
		return new(RefundRequest)
	},
}

// GetRefundRequest() 从对象池中获取RefundRequest
func GetRefundRequest() *RefundRequest {
	return poolRefundRequest.Get().(*RefundRequest)
}

// ReleaseRefundRequest 释放RefundRequest
func ReleaseRefundRequest(v *RefundRequest) {
	v.OpId = ""
	v.Operator = ""
	v.OriginalOpId = ""
	v.OuterStoreId = ""
	v.PayCode = ""
	v.SaleTicket = ""
	v.Amount = 0
	poolRefundRequest.Put(v)
}
