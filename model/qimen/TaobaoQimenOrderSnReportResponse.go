package qimen

import (
	"sync"
)

// TaobaoQimenOrderSnReportResponse 结构体
type TaobaoQimenOrderSnReportResponse struct {
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolTaobaoQimenOrderSnReportResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenOrderSnReportResponse)
	},
}

// GetTaobaoQimenOrderSnReportResponse() 从对象池中获取TaobaoQimenOrderSnReportResponse
func GetTaobaoQimenOrderSnReportResponse() *TaobaoQimenOrderSnReportResponse {
	return poolTaobaoQimenOrderSnReportResponse.Get().(*TaobaoQimenOrderSnReportResponse)
}

// ReleaseTaobaoQimenOrderSnReportResponse 释放TaobaoQimenOrderSnReportResponse
func ReleaseTaobaoQimenOrderSnReportResponse(v *TaobaoQimenOrderSnReportResponse) {
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	poolTaobaoQimenOrderSnReportResponse.Put(v)
}
