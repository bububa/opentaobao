package qimen

import (
	"sync"
)

// TaobaoQimenCombineitemQueryResponse 结构体
type TaobaoQimenCombineitemQueryResponse struct {
	// 奇门仓储字段
	CombItems []CombItem `json:"combItems,omitempty" xml:"combItems>comb_item,omitempty"`
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolTaobaoQimenCombineitemQueryResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenCombineitemQueryResponse)
	},
}

// GetTaobaoQimenCombineitemQueryResponse() 从对象池中获取TaobaoQimenCombineitemQueryResponse
func GetTaobaoQimenCombineitemQueryResponse() *TaobaoQimenCombineitemQueryResponse {
	return poolTaobaoQimenCombineitemQueryResponse.Get().(*TaobaoQimenCombineitemQueryResponse)
}

// ReleaseTaobaoQimenCombineitemQueryResponse 释放TaobaoQimenCombineitemQueryResponse
func ReleaseTaobaoQimenCombineitemQueryResponse(v *TaobaoQimenCombineitemQueryResponse) {
	v.CombItems = v.CombItems[:0]
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	poolTaobaoQimenCombineitemQueryResponse.Put(v)
}
