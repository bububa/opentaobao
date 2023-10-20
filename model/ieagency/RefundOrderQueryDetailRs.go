package ieagency

import (
	"sync"
)

// RefundOrderQueryDetailRs 结构体
type RefundOrderQueryDetailRs struct {
	// 错误描述
	ApiErrorMsg string `json:"api_error_msg,omitempty" xml:"api_error_msg,omitempty"`
	// 错误trace
	ErrTrace string `json:"err_trace,omitempty" xml:"err_trace,omitempty"`
	// 机器名称
	HostName string `json:"host_name,omitempty" xml:"host_name,omitempty"`
	// 错误码
	ApiErrorCode int64 `json:"api_error_code,omitempty" xml:"api_error_code,omitempty"`
	// 申请单详情
	RefundOrderVo *RefundOrderVo `json:"refund_order_vo,omitempty" xml:"refund_order_vo,omitempty"`
	// 失败
	Failure bool `json:"failure,omitempty" xml:"failure,omitempty"`
	// 成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolRefundOrderQueryDetailRs = sync.Pool{
	New: func() any {
		return new(RefundOrderQueryDetailRs)
	},
}

// GetRefundOrderQueryDetailRs() 从对象池中获取RefundOrderQueryDetailRs
func GetRefundOrderQueryDetailRs() *RefundOrderQueryDetailRs {
	return poolRefundOrderQueryDetailRs.Get().(*RefundOrderQueryDetailRs)
}

// ReleaseRefundOrderQueryDetailRs 释放RefundOrderQueryDetailRs
func ReleaseRefundOrderQueryDetailRs(v *RefundOrderQueryDetailRs) {
	v.ApiErrorMsg = ""
	v.ErrTrace = ""
	v.HostName = ""
	v.ApiErrorCode = 0
	v.RefundOrderVo = nil
	v.Failure = false
	v.Success = false
	poolRefundOrderQueryDetailRs.Put(v)
}
