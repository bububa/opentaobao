package tuike

import (
	"sync"
)

// AlibabaTuikeOfferZhitokenResult 结构体
type AlibabaTuikeOfferZhitokenResult struct {
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

var poolAlibabaTuikeOfferZhitokenResult = sync.Pool{
	New: func() any {
		return new(AlibabaTuikeOfferZhitokenResult)
	},
}

// GetAlibabaTuikeOfferZhitokenResult() 从对象池中获取AlibabaTuikeOfferZhitokenResult
func GetAlibabaTuikeOfferZhitokenResult() *AlibabaTuikeOfferZhitokenResult {
	return poolAlibabaTuikeOfferZhitokenResult.Get().(*AlibabaTuikeOfferZhitokenResult)
}

// ReleaseAlibabaTuikeOfferZhitokenResult 释放AlibabaTuikeOfferZhitokenResult
func ReleaseAlibabaTuikeOfferZhitokenResult(v *AlibabaTuikeOfferZhitokenResult) {
	v.Headers = ""
	v.Model = ""
	v.MsgCode = ""
	v.BizExtMap = ""
	v.MsgInfo = ""
	v.HttpStatusCode = 0
	v.Success = false
	poolAlibabaTuikeOfferZhitokenResult.Put(v)
}
