package qimen

import (
	"sync"
)

// TaobaoQimenInventoryReportResponse 结构体
type TaobaoQimenInventoryReportResponse struct {
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolTaobaoQimenInventoryReportResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenInventoryReportResponse)
	},
}

// GetTaobaoQimenInventoryReportResponse() 从对象池中获取TaobaoQimenInventoryReportResponse
func GetTaobaoQimenInventoryReportResponse() *TaobaoQimenInventoryReportResponse {
	return poolTaobaoQimenInventoryReportResponse.Get().(*TaobaoQimenInventoryReportResponse)
}

// ReleaseTaobaoQimenInventoryReportResponse 释放TaobaoQimenInventoryReportResponse
func ReleaseTaobaoQimenInventoryReportResponse(v *TaobaoQimenInventoryReportResponse) {
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	poolTaobaoQimenInventoryReportResponse.Put(v)
}
