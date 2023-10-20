package jae

import (
	"sync"
)

// TaobaoAplatformWeakgetResult 结构体
type TaobaoAplatformWeakgetResult struct {
	// headers
	Headers string `json:"headers,omitempty" xml:"headers,omitempty"`
	// model
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// bizExtMap
	BizExtMap string `json:"biz_ext_map,omitempty" xml:"biz_ext_map,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// httpStatusCode
	HttpStatusCode int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoAplatformWeakgetResult = sync.Pool{
	New: func() any {
		return new(TaobaoAplatformWeakgetResult)
	},
}

// GetTaobaoAplatformWeakgetResult() 从对象池中获取TaobaoAplatformWeakgetResult
func GetTaobaoAplatformWeakgetResult() *TaobaoAplatformWeakgetResult {
	return poolTaobaoAplatformWeakgetResult.Get().(*TaobaoAplatformWeakgetResult)
}

// ReleaseTaobaoAplatformWeakgetResult 释放TaobaoAplatformWeakgetResult
func ReleaseTaobaoAplatformWeakgetResult(v *TaobaoAplatformWeakgetResult) {
	v.Headers = ""
	v.Model = ""
	v.MsgCode = ""
	v.BizExtMap = ""
	v.MsgInfo = ""
	v.HttpStatusCode = 0
	v.Success = false
	poolTaobaoAplatformWeakgetResult.Put(v)
}
