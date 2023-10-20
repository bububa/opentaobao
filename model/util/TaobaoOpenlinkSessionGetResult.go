package util

import (
	"sync"
)

// TaobaoOpenlinkSessionGetResult 结构体
type TaobaoOpenlinkSessionGetResult struct {
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// model
	Model *SessionInfo `json:"model,omitempty" xml:"model,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoOpenlinkSessionGetResult = sync.Pool{
	New: func() any {
		return new(TaobaoOpenlinkSessionGetResult)
	},
}

// GetTaobaoOpenlinkSessionGetResult() 从对象池中获取TaobaoOpenlinkSessionGetResult
func GetTaobaoOpenlinkSessionGetResult() *TaobaoOpenlinkSessionGetResult {
	return poolTaobaoOpenlinkSessionGetResult.Get().(*TaobaoOpenlinkSessionGetResult)
}

// ReleaseTaobaoOpenlinkSessionGetResult 释放TaobaoOpenlinkSessionGetResult
func ReleaseTaobaoOpenlinkSessionGetResult(v *TaobaoOpenlinkSessionGetResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.Success = false
	poolTaobaoOpenlinkSessionGetResult.Put(v)
}
