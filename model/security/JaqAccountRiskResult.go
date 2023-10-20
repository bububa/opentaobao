package security

import (
	"sync"
)

// JaqAccountRiskResult 结构体
type JaqAccountRiskResult struct {
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 风控结果的真正数据内容
	Data *JaqAccountRiskData `json:"data,omitempty" xml:"data,omitempty"`
	// 错误码
	ErrorCode int64 `json:"error_code,omitempty" xml:"error_code,omitempty"`
}

var poolJaqAccountRiskResult = sync.Pool{
	New: func() any {
		return new(JaqAccountRiskResult)
	},
}

// GetJaqAccountRiskResult() 从对象池中获取JaqAccountRiskResult
func GetJaqAccountRiskResult() *JaqAccountRiskResult {
	return poolJaqAccountRiskResult.Get().(*JaqAccountRiskResult)
}

// ReleaseJaqAccountRiskResult 释放JaqAccountRiskResult
func ReleaseJaqAccountRiskResult(v *JaqAccountRiskResult) {
	v.ErrorMsg = ""
	v.Data = nil
	v.ErrorCode = 0
	poolJaqAccountRiskResult.Put(v)
}
