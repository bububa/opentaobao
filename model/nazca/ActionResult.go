package nazca

import (
	"sync"
)

// ActionResult 结构体
type ActionResult struct {
	// error
	Error string `json:"error,omitempty" xml:"error,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 错误信息
	SubErrorMessage string `json:"sub_error_message,omitempty" xml:"sub_error_message,omitempty"`
	// retValue
	RetValue *AuthApplyDoneCallBackDo `json:"ret_value,omitempty" xml:"ret_value,omitempty"`
	// 成功状态
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolActionResult = sync.Pool{
	New: func() any {
		return new(ActionResult)
	},
}

// GetActionResult() 从对象池中获取ActionResult
func GetActionResult() *ActionResult {
	return poolActionResult.Get().(*ActionResult)
}

// ReleaseActionResult 释放ActionResult
func ReleaseActionResult(v *ActionResult) {
	v.Error = ""
	v.Message = ""
	v.SubErrorMessage = ""
	v.RetValue = nil
	v.Success = false
	poolActionResult.Put(v)
}
