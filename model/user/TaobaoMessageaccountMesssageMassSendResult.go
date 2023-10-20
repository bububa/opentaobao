package user

import (
	"sync"
)

// TaobaoMessageaccountMesssageMassSendResult 结构体
type TaobaoMessageaccountMesssageMassSendResult struct {
	// model
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// errMessage
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// errCode
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// success
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

var poolTaobaoMessageaccountMesssageMassSendResult = sync.Pool{
	New: func() any {
		return new(TaobaoMessageaccountMesssageMassSendResult)
	},
}

// GetTaobaoMessageaccountMesssageMassSendResult() 从对象池中获取TaobaoMessageaccountMesssageMassSendResult
func GetTaobaoMessageaccountMesssageMassSendResult() *TaobaoMessageaccountMesssageMassSendResult {
	return poolTaobaoMessageaccountMesssageMassSendResult.Get().(*TaobaoMessageaccountMesssageMassSendResult)
}

// ReleaseTaobaoMessageaccountMesssageMassSendResult 释放TaobaoMessageaccountMesssageMassSendResult
func ReleaseTaobaoMessageaccountMesssageMassSendResult(v *TaobaoMessageaccountMesssageMassSendResult) {
	v.Model = ""
	v.ErrMessage = ""
	v.ErrCode = ""
	v.Result = false
	poolTaobaoMessageaccountMesssageMassSendResult.Put(v)
}
