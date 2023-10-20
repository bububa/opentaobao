package user

import (
	"sync"
)

// TaobaoMiniappMesssageReplyResult 结构体
type TaobaoMiniappMesssageReplyResult struct {
	// model
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// errMessage
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// errCode
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoMiniappMesssageReplyResult = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappMesssageReplyResult)
	},
}

// GetTaobaoMiniappMesssageReplyResult() 从对象池中获取TaobaoMiniappMesssageReplyResult
func GetTaobaoMiniappMesssageReplyResult() *TaobaoMiniappMesssageReplyResult {
	return poolTaobaoMiniappMesssageReplyResult.Get().(*TaobaoMiniappMesssageReplyResult)
}

// ReleaseTaobaoMiniappMesssageReplyResult 释放TaobaoMiniappMesssageReplyResult
func ReleaseTaobaoMiniappMesssageReplyResult(v *TaobaoMiniappMesssageReplyResult) {
	v.Model = ""
	v.ErrMessage = ""
	v.ErrCode = ""
	v.Success = false
	poolTaobaoMiniappMesssageReplyResult.Put(v)
}
