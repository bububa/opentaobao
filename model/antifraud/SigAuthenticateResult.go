package antifraud

import (
	"sync"
)

// SigAuthenticateResult 结构体
type SigAuthenticateResult struct {
	// 100	验证通过	验证通过 200	服务器故障,此时ResultWrapper的success=false	服务自身正确识别的服务器故障行为，请视同验证通过处理  900	验证不通过	预留9XX做为扩展，901:NOPASS_USER_APP，应用方传入参数有误,如appkey与access_key不匹配等
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误描述消息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// detail
	Detail string `json:"detail,omitempty" xml:"detail,omitempty"`
	// 签名串生成的毫秒值(System.currentTimeMillis()).使用方自行判断此签名串是否已过期
	Timestamp int64 `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

var poolSigAuthenticateResult = sync.Pool{
	New: func() any {
		return new(SigAuthenticateResult)
	},
}

// GetSigAuthenticateResult() 从对象池中获取SigAuthenticateResult
func GetSigAuthenticateResult() *SigAuthenticateResult {
	return poolSigAuthenticateResult.Get().(*SigAuthenticateResult)
}

// ReleaseSigAuthenticateResult 释放SigAuthenticateResult
func ReleaseSigAuthenticateResult(v *SigAuthenticateResult) {
	v.Code = ""
	v.Msg = ""
	v.Detail = ""
	v.Timestamp = 0
	poolSigAuthenticateResult.Put(v)
}
