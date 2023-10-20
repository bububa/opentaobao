package flight

import (
	"sync"
)

// RefundOrderFillConfirmFeeRs 结构体
type RefundOrderFillConfirmFeeRs struct {
	// 错误秒速
	ApiErrorMsg string `json:"api_error_msg,omitempty" xml:"api_error_msg,omitempty"`
	// errTrace
	ErrTrace string `json:"err_trace,omitempty" xml:"err_trace,omitempty"`
	// 主机名称
	HostName string `json:"host_name,omitempty" xml:"host_name,omitempty"`
	// 错误编码
	ApiErrorCode int64 `json:"api_error_code,omitempty" xml:"api_error_code,omitempty"`
	// 是否失败
	Failure bool `json:"failure,omitempty" xml:"failure,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolRefundOrderFillConfirmFeeRs = sync.Pool{
	New: func() any {
		return new(RefundOrderFillConfirmFeeRs)
	},
}

// GetRefundOrderFillConfirmFeeRs() 从对象池中获取RefundOrderFillConfirmFeeRs
func GetRefundOrderFillConfirmFeeRs() *RefundOrderFillConfirmFeeRs {
	return poolRefundOrderFillConfirmFeeRs.Get().(*RefundOrderFillConfirmFeeRs)
}

// ReleaseRefundOrderFillConfirmFeeRs 释放RefundOrderFillConfirmFeeRs
func ReleaseRefundOrderFillConfirmFeeRs(v *RefundOrderFillConfirmFeeRs) {
	v.ApiErrorMsg = ""
	v.ErrTrace = ""
	v.HostName = ""
	v.ApiErrorCode = 0
	v.Failure = false
	v.Success = false
	poolRefundOrderFillConfirmFeeRs.Put(v)
}
