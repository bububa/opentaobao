package uscesl

import (
	"sync"
)

// TaobaoUsceslBizApActivateResult 结构体
type TaobaoUsceslBizApActivateResult struct {
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误编码
	BusinessCode string `json:"business_code,omitempty" xml:"business_code,omitempty"`
	// 1标识成功,其他失败
	ReturnCode int64 `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 业务true或false
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 请求true或false
	Target bool `json:"target,omitempty" xml:"target,omitempty"`
}

var poolTaobaoUsceslBizApActivateResult = sync.Pool{
	New: func() any {
		return new(TaobaoUsceslBizApActivateResult)
	},
}

// GetTaobaoUsceslBizApActivateResult() 从对象池中获取TaobaoUsceslBizApActivateResult
func GetTaobaoUsceslBizApActivateResult() *TaobaoUsceslBizApActivateResult {
	return poolTaobaoUsceslBizApActivateResult.Get().(*TaobaoUsceslBizApActivateResult)
}

// ReleaseTaobaoUsceslBizApActivateResult 释放TaobaoUsceslBizApActivateResult
func ReleaseTaobaoUsceslBizApActivateResult(v *TaobaoUsceslBizApActivateResult) {
	v.Message = ""
	v.BusinessCode = ""
	v.ReturnCode = 0
	v.IsSuccess = false
	v.Target = false
	poolTaobaoUsceslBizApActivateResult.Put(v)
}
