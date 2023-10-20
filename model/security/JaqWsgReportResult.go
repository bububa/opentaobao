package security

import (
	"sync"
)

// JaqWsgReportResult 结构体
type JaqWsgReportResult struct {
	// 错误描述
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 生成的token
	Token string `json:"token,omitempty" xml:"token,omitempty"`
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 安全验证前向化下发参数结构体
	JaqDispatchParam *JaqDispatchParam `json:"jaq_dispatch_param,omitempty" xml:"jaq_dispatch_param,omitempty"`
}

var poolJaqWsgReportResult = sync.Pool{
	New: func() any {
		return new(JaqWsgReportResult)
	},
}

// GetJaqWsgReportResult() 从对象池中获取JaqWsgReportResult
func GetJaqWsgReportResult() *JaqWsgReportResult {
	return poolJaqWsgReportResult.Get().(*JaqWsgReportResult)
}

// ReleaseJaqWsgReportResult 释放JaqWsgReportResult
func ReleaseJaqWsgReportResult(v *JaqWsgReportResult) {
	v.ErrorMsg = ""
	v.Token = ""
	v.ErrorCode = 0
	v.JaqDispatchParam = nil
	poolJaqWsgReportResult.Put(v)
}
