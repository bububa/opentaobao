package qimen

import (
	"sync"
)

// TaobaoQimenExpressinfoQueryResponse 结构体
type TaobaoQimenExpressinfoQueryResponse struct {
	// 奇门仓储字段
	ExpressInfos []ExpressInfo `json:"expressInfos,omitempty" xml:"expressInfos>express_info,omitempty"`
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolTaobaoQimenExpressinfoQueryResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenExpressinfoQueryResponse)
	},
}

// GetTaobaoQimenExpressinfoQueryResponse() 从对象池中获取TaobaoQimenExpressinfoQueryResponse
func GetTaobaoQimenExpressinfoQueryResponse() *TaobaoQimenExpressinfoQueryResponse {
	return poolTaobaoQimenExpressinfoQueryResponse.Get().(*TaobaoQimenExpressinfoQueryResponse)
}

// ReleaseTaobaoQimenExpressinfoQueryResponse 释放TaobaoQimenExpressinfoQueryResponse
func ReleaseTaobaoQimenExpressinfoQueryResponse(v *TaobaoQimenExpressinfoQueryResponse) {
	v.ExpressInfos = v.ExpressInfos[:0]
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	poolTaobaoQimenExpressinfoQueryResponse.Put(v)
}
