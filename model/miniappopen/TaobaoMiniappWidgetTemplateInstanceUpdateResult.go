package miniappopen

import (
	"sync"
)

// TaobaoMiniappWidgetTemplateInstanceUpdateResult 结构体
type TaobaoMiniappWidgetTemplateInstanceUpdateResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// 返回结果
	Success string `json:"success,omitempty" xml:"success,omitempty"`
	// 实体信息
	Model *MiniAppEntityTemplateDto `json:"model,omitempty" xml:"model,omitempty"`
}

var poolTaobaoMiniappWidgetTemplateInstanceUpdateResult = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappWidgetTemplateInstanceUpdateResult)
	},
}

// GetTaobaoMiniappWidgetTemplateInstanceUpdateResult() 从对象池中获取TaobaoMiniappWidgetTemplateInstanceUpdateResult
func GetTaobaoMiniappWidgetTemplateInstanceUpdateResult() *TaobaoMiniappWidgetTemplateInstanceUpdateResult {
	return poolTaobaoMiniappWidgetTemplateInstanceUpdateResult.Get().(*TaobaoMiniappWidgetTemplateInstanceUpdateResult)
}

// ReleaseTaobaoMiniappWidgetTemplateInstanceUpdateResult 释放TaobaoMiniappWidgetTemplateInstanceUpdateResult
func ReleaseTaobaoMiniappWidgetTemplateInstanceUpdateResult(v *TaobaoMiniappWidgetTemplateInstanceUpdateResult) {
	v.ErrCode = ""
	v.ErrMessage = ""
	v.Success = ""
	v.Model = nil
	poolTaobaoMiniappWidgetTemplateInstanceUpdateResult.Put(v)
}
