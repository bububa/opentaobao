package uscesl

import (
	"sync"
)

// TaobaoUsceslBizItemLightUpResult 结构体
type TaobaoUsceslBizItemLightUpResult struct {
	// 返回码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 业务错误码
	BusinessCode string `json:"business_code,omitempty" xml:"business_code,omitempty"`
	// true或false
	Target *LightResultInfoBo `json:"target,omitempty" xml:"target,omitempty"`
	// true或false
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolTaobaoUsceslBizItemLightUpResult = sync.Pool{
	New: func() any {
		return new(TaobaoUsceslBizItemLightUpResult)
	},
}

// GetTaobaoUsceslBizItemLightUpResult() 从对象池中获取TaobaoUsceslBizItemLightUpResult
func GetTaobaoUsceslBizItemLightUpResult() *TaobaoUsceslBizItemLightUpResult {
	return poolTaobaoUsceslBizItemLightUpResult.Get().(*TaobaoUsceslBizItemLightUpResult)
}

// ReleaseTaobaoUsceslBizItemLightUpResult 释放TaobaoUsceslBizItemLightUpResult
func ReleaseTaobaoUsceslBizItemLightUpResult(v *TaobaoUsceslBizItemLightUpResult) {
	v.ReturnCode = ""
	v.Message = ""
	v.BusinessCode = ""
	v.Target = nil
	v.IsSuccess = false
	poolTaobaoUsceslBizItemLightUpResult.Put(v)
}
