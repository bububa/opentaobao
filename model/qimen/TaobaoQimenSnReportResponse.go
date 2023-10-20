package qimen

import (
	"sync"
)

// TaobaoQimenSnReportResponse 结构体
type TaobaoQimenSnReportResponse struct {
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolTaobaoQimenSnReportResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenSnReportResponse)
	},
}

// GetTaobaoQimenSnReportResponse() 从对象池中获取TaobaoQimenSnReportResponse
func GetTaobaoQimenSnReportResponse() *TaobaoQimenSnReportResponse {
	return poolTaobaoQimenSnReportResponse.Get().(*TaobaoQimenSnReportResponse)
}

// ReleaseTaobaoQimenSnReportResponse 释放TaobaoQimenSnReportResponse
func ReleaseTaobaoQimenSnReportResponse(v *TaobaoQimenSnReportResponse) {
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	poolTaobaoQimenSnReportResponse.Put(v)
}
