package miniappopen

import (
	"sync"
)

// TaobaoMiniapppTemplateInstantiateResult 结构体
type TaobaoMiniapppTemplateInstantiateResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// model
	Model *MiniAppEntityDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoMiniapppTemplateInstantiateResult = sync.Pool{
	New: func() any {
		return new(TaobaoMiniapppTemplateInstantiateResult)
	},
}

// GetTaobaoMiniapppTemplateInstantiateResult() 从对象池中获取TaobaoMiniapppTemplateInstantiateResult
func GetTaobaoMiniapppTemplateInstantiateResult() *TaobaoMiniapppTemplateInstantiateResult {
	return poolTaobaoMiniapppTemplateInstantiateResult.Get().(*TaobaoMiniapppTemplateInstantiateResult)
}

// ReleaseTaobaoMiniapppTemplateInstantiateResult 释放TaobaoMiniapppTemplateInstantiateResult
func ReleaseTaobaoMiniapppTemplateInstantiateResult(v *TaobaoMiniapppTemplateInstantiateResult) {
	v.ErrCode = ""
	v.ErrMessage = ""
	v.Model = nil
	v.Success = false
	poolTaobaoMiniapppTemplateInstantiateResult.Put(v)
}
