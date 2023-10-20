package eticket

import (
	"sync"
)

// QueryTbMaCallbackReq 结构体
type QueryTbMaCallbackReq struct {
	// 淘宝码值
	Code string `json:"code,omitempty" xml:"code,omitempty"`
}

var poolQueryTbMaCallbackReq = sync.Pool{
	New: func() any {
		return new(QueryTbMaCallbackReq)
	},
}

// GetQueryTbMaCallbackReq() 从对象池中获取QueryTbMaCallbackReq
func GetQueryTbMaCallbackReq() *QueryTbMaCallbackReq {
	return poolQueryTbMaCallbackReq.Get().(*QueryTbMaCallbackReq)
}

// ReleaseQueryTbMaCallbackReq 释放QueryTbMaCallbackReq
func ReleaseQueryTbMaCallbackReq(v *QueryTbMaCallbackReq) {
	v.Code = ""
	poolQueryTbMaCallbackReq.Put(v)
}
