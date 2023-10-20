package ieagency

import (
	"sync"
)

// RefundMoneyNoPasswordRs 结构体
type RefundMoneyNoPasswordRs struct {
	// apiErrorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// apiErrorCode
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolRefundMoneyNoPasswordRs = sync.Pool{
	New: func() any {
		return new(RefundMoneyNoPasswordRs)
	},
}

// GetRefundMoneyNoPasswordRs() 从对象池中获取RefundMoneyNoPasswordRs
func GetRefundMoneyNoPasswordRs() *RefundMoneyNoPasswordRs {
	return poolRefundMoneyNoPasswordRs.Get().(*RefundMoneyNoPasswordRs)
}

// ReleaseRefundMoneyNoPasswordRs 释放RefundMoneyNoPasswordRs
func ReleaseRefundMoneyNoPasswordRs(v *RefundMoneyNoPasswordRs) {
	v.ErrorMsg = ""
	v.ErrorCode = 0
	v.Success = false
	poolRefundMoneyNoPasswordRs.Put(v)
}
