package user

import (
	"sync"
)

// TaobaoMessageaccountMesssageReplyResult 结构体
type TaobaoMessageaccountMesssageReplyResult struct {
	// model
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// errMessage
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// errCode
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// success
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

var poolTaobaoMessageaccountMesssageReplyResult = sync.Pool{
	New: func() any {
		return new(TaobaoMessageaccountMesssageReplyResult)
	},
}

// GetTaobaoMessageaccountMesssageReplyResult() 从对象池中获取TaobaoMessageaccountMesssageReplyResult
func GetTaobaoMessageaccountMesssageReplyResult() *TaobaoMessageaccountMesssageReplyResult {
	return poolTaobaoMessageaccountMesssageReplyResult.Get().(*TaobaoMessageaccountMesssageReplyResult)
}

// ReleaseTaobaoMessageaccountMesssageReplyResult 释放TaobaoMessageaccountMesssageReplyResult
func ReleaseTaobaoMessageaccountMesssageReplyResult(v *TaobaoMessageaccountMesssageReplyResult) {
	v.Model = ""
	v.ErrMessage = ""
	v.ErrCode = ""
	v.Result = false
	poolTaobaoMessageaccountMesssageReplyResult.Put(v)
}
