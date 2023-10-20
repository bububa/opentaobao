package qimen

import (
	"sync"
)

// TaobaoQimenEntryorderCreateResponse 结构体
type TaobaoQimenEntryorderCreateResponse struct {
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 仓储系统入库单编码
	EntryOrderId string `json:"entryOrderId,omitempty" xml:"entryOrderId,omitempty"`
}

var poolTaobaoQimenEntryorderCreateResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenEntryorderCreateResponse)
	},
}

// GetTaobaoQimenEntryorderCreateResponse() 从对象池中获取TaobaoQimenEntryorderCreateResponse
func GetTaobaoQimenEntryorderCreateResponse() *TaobaoQimenEntryorderCreateResponse {
	return poolTaobaoQimenEntryorderCreateResponse.Get().(*TaobaoQimenEntryorderCreateResponse)
}

// ReleaseTaobaoQimenEntryorderCreateResponse 释放TaobaoQimenEntryorderCreateResponse
func ReleaseTaobaoQimenEntryorderCreateResponse(v *TaobaoQimenEntryorderCreateResponse) {
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	v.EntryOrderId = ""
	poolTaobaoQimenEntryorderCreateResponse.Put(v)
}
