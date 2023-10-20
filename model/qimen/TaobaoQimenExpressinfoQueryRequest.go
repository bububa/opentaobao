package qimen

import (
	"sync"
)

// TaobaoQimenExpressinfoQueryRequest 结构体
type TaobaoQimenExpressinfoQueryRequest struct {
	// 奇门仓储字段
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 奇门仓储字段
	ExpressCode string `json:"expressCode,omitempty" xml:"expressCode,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenExpressinfoQueryMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}

var poolTaobaoQimenExpressinfoQueryRequest = sync.Pool{
	New: func() any {
		return new(TaobaoQimenExpressinfoQueryRequest)
	},
}

// GetTaobaoQimenExpressinfoQueryRequest() 从对象池中获取TaobaoQimenExpressinfoQueryRequest
func GetTaobaoQimenExpressinfoQueryRequest() *TaobaoQimenExpressinfoQueryRequest {
	return poolTaobaoQimenExpressinfoQueryRequest.Get().(*TaobaoQimenExpressinfoQueryRequest)
}

// ReleaseTaobaoQimenExpressinfoQueryRequest 释放TaobaoQimenExpressinfoQueryRequest
func ReleaseTaobaoQimenExpressinfoQueryRequest(v *TaobaoQimenExpressinfoQueryRequest) {
	v.OwnerCode = ""
	v.ExpressCode = ""
	v.ExtendProps = nil
	poolTaobaoQimenExpressinfoQueryRequest.Put(v)
}
