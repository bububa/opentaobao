package user

import (
	"sync"
)

// TaobaoMiniappUserInfoGetResult 结构体
type TaobaoMiniappUserInfoGetResult struct {
	// 错误信息
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// model
	Model *OpenUserInfoDto `json:"model,omitempty" xml:"model,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoMiniappUserInfoGetResult = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappUserInfoGetResult)
	},
}

// GetTaobaoMiniappUserInfoGetResult() 从对象池中获取TaobaoMiniappUserInfoGetResult
func GetTaobaoMiniappUserInfoGetResult() *TaobaoMiniappUserInfoGetResult {
	return poolTaobaoMiniappUserInfoGetResult.Get().(*TaobaoMiniappUserInfoGetResult)
}

// ReleaseTaobaoMiniappUserInfoGetResult 释放TaobaoMiniappUserInfoGetResult
func ReleaseTaobaoMiniappUserInfoGetResult(v *TaobaoMiniappUserInfoGetResult) {
	v.ErrMessage = ""
	v.ErrCode = ""
	v.Model = nil
	v.Success = false
	poolTaobaoMiniappUserInfoGetResult.Put(v)
}
