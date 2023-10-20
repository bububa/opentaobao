package miniappopen

import (
	"sync"
)

// TaobaoMiniappTemplateUpdateResult 结构体
type TaobaoMiniappTemplateUpdateResult struct {
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// model
	Model *MiniAppEntityTemplateDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoMiniappTemplateUpdateResult = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappTemplateUpdateResult)
	},
}

// GetTaobaoMiniappTemplateUpdateResult() 从对象池中获取TaobaoMiniappTemplateUpdateResult
func GetTaobaoMiniappTemplateUpdateResult() *TaobaoMiniappTemplateUpdateResult {
	return poolTaobaoMiniappTemplateUpdateResult.Get().(*TaobaoMiniappTemplateUpdateResult)
}

// ReleaseTaobaoMiniappTemplateUpdateResult 释放TaobaoMiniappTemplateUpdateResult
func ReleaseTaobaoMiniappTemplateUpdateResult(v *TaobaoMiniappTemplateUpdateResult) {
	v.ErrCode = ""
	v.ErrMessage = ""
	v.Model = nil
	v.Success = false
	poolTaobaoMiniappTemplateUpdateResult.Put(v)
}
