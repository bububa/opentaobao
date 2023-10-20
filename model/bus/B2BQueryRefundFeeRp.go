package bus

import (
	"sync"
)

// B2BQueryRefundFeeRp 结构体
type B2BQueryRefundFeeRp struct {
	// refundFees
	RefundFees []string `json:"refund_fees,omitempty" xml:"refund_fees>string,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// canReturnSingleTicket
	CanReturnSingleTicket bool `json:"can_return_single_ticket,omitempty" xml:"can_return_single_ticket,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolB2BQueryRefundFeeRp = sync.Pool{
	New: func() any {
		return new(B2BQueryRefundFeeRp)
	},
}

// GetB2BQueryRefundFeeRp() 从对象池中获取B2BQueryRefundFeeRp
func GetB2BQueryRefundFeeRp() *B2BQueryRefundFeeRp {
	return poolB2BQueryRefundFeeRp.Get().(*B2BQueryRefundFeeRp)
}

// ReleaseB2BQueryRefundFeeRp 释放B2BQueryRefundFeeRp
func ReleaseB2BQueryRefundFeeRp(v *B2BQueryRefundFeeRp) {
	v.RefundFees = v.RefundFees[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.CanReturnSingleTicket = false
	v.Success = false
	poolB2BQueryRefundFeeRp.Put(v)
}
