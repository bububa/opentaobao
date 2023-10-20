package tbrefund

import (
	"sync"
)

// RefundMappingResult 结构体
type RefundMappingResult struct {
	// 结果信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 退款ID
	RefundId string `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
	// 是否成功
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
}

var poolRefundMappingResult = sync.Pool{
	New: func() any {
		return new(RefundMappingResult)
	},
}

// GetRefundMappingResult() 从对象池中获取RefundMappingResult
func GetRefundMappingResult() *RefundMappingResult {
	return poolRefundMappingResult.Get().(*RefundMappingResult)
}

// ReleaseRefundMappingResult 释放RefundMappingResult
func ReleaseRefundMappingResult(v *RefundMappingResult) {
	v.Message = ""
	v.RefundId = ""
	v.Succ = false
	poolRefundMappingResult.Put(v)
}
