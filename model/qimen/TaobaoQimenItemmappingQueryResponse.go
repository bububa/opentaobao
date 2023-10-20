package qimen

import (
	"sync"
)

// TaobaoQimenItemmappingQueryResponse 结构体
type TaobaoQimenItemmappingQueryResponse struct {
	// 商品映射关系
	ItemMappings []ItemMapping `json:"itemMappings,omitempty" xml:"itemMappings>item_mapping,omitempty"`
	// 响应结果:success|failure,success,string(10),必填
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码,0,string(50),
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息,invalid appkey,string(100),
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolTaobaoQimenItemmappingQueryResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenItemmappingQueryResponse)
	},
}

// GetTaobaoQimenItemmappingQueryResponse() 从对象池中获取TaobaoQimenItemmappingQueryResponse
func GetTaobaoQimenItemmappingQueryResponse() *TaobaoQimenItemmappingQueryResponse {
	return poolTaobaoQimenItemmappingQueryResponse.Get().(*TaobaoQimenItemmappingQueryResponse)
}

// ReleaseTaobaoQimenItemmappingQueryResponse 释放TaobaoQimenItemmappingQueryResponse
func ReleaseTaobaoQimenItemmappingQueryResponse(v *TaobaoQimenItemmappingQueryResponse) {
	v.ItemMappings = v.ItemMappings[:0]
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	poolTaobaoQimenItemmappingQueryResponse.Put(v)
}
