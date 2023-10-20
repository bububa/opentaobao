package qimen

import (
	"sync"
)

// TaobaoQimenInventorybatchQueryResponse 结构体
type TaobaoQimenInventorybatchQueryResponse struct {
	// success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 总行数，必填
	TotalCount int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// 明细列表
	Items *Items `json:"items,omitempty" xml:"items,omitempty"`
}

var poolTaobaoQimenInventorybatchQueryResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenInventorybatchQueryResponse)
	},
}

// GetTaobaoQimenInventorybatchQueryResponse() 从对象池中获取TaobaoQimenInventorybatchQueryResponse
func GetTaobaoQimenInventorybatchQueryResponse() *TaobaoQimenInventorybatchQueryResponse {
	return poolTaobaoQimenInventorybatchQueryResponse.Get().(*TaobaoQimenInventorybatchQueryResponse)
}

// ReleaseTaobaoQimenInventorybatchQueryResponse 释放TaobaoQimenInventorybatchQueryResponse
func ReleaseTaobaoQimenInventorybatchQueryResponse(v *TaobaoQimenInventorybatchQueryResponse) {
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	v.TotalCount = 0
	v.Items = nil
	poolTaobaoQimenInventorybatchQueryResponse.Put(v)
}
