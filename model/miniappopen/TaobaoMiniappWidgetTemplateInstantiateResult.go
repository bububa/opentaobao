package miniappopen

import (
	"sync"
)

// TaobaoMiniappWidgetTemplateInstantiateResult 结构体
type TaobaoMiniappWidgetTemplateInstantiateResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// 返回实体信息
	Model *MiniAppEntityTemplateDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoMiniappWidgetTemplateInstantiateResult = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappWidgetTemplateInstantiateResult)
	},
}

// GetTaobaoMiniappWidgetTemplateInstantiateResult() 从对象池中获取TaobaoMiniappWidgetTemplateInstantiateResult
func GetTaobaoMiniappWidgetTemplateInstantiateResult() *TaobaoMiniappWidgetTemplateInstantiateResult {
	return poolTaobaoMiniappWidgetTemplateInstantiateResult.Get().(*TaobaoMiniappWidgetTemplateInstantiateResult)
}

// ReleaseTaobaoMiniappWidgetTemplateInstantiateResult 释放TaobaoMiniappWidgetTemplateInstantiateResult
func ReleaseTaobaoMiniappWidgetTemplateInstantiateResult(v *TaobaoMiniappWidgetTemplateInstantiateResult) {
	v.ErrCode = ""
	v.ErrMessage = ""
	v.Model = nil
	v.Success = false
	poolTaobaoMiniappWidgetTemplateInstantiateResult.Put(v)
}
