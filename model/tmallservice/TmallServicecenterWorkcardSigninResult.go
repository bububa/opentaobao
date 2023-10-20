package tmallservice

import (
	"sync"
)

// TmallServicecenterWorkcardSigninResult 结构体
type TmallServicecenterWorkcardSigninResult struct {
	// 错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 请求是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallServicecenterWorkcardSigninResult = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardSigninResult)
	},
}

// GetTmallServicecenterWorkcardSigninResult() 从对象池中获取TmallServicecenterWorkcardSigninResult
func GetTmallServicecenterWorkcardSigninResult() *TmallServicecenterWorkcardSigninResult {
	return poolTmallServicecenterWorkcardSigninResult.Get().(*TmallServicecenterWorkcardSigninResult)
}

// ReleaseTmallServicecenterWorkcardSigninResult 释放TmallServicecenterWorkcardSigninResult
func ReleaseTmallServicecenterWorkcardSigninResult(v *TmallServicecenterWorkcardSigninResult) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Success = false
	poolTmallServicecenterWorkcardSigninResult.Put(v)
}
