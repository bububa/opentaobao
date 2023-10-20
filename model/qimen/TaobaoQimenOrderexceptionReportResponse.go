package qimen

import (
	"sync"
)

// TaobaoQimenOrderexceptionReportResponse 结构体
type TaobaoQimenOrderexceptionReportResponse struct {
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolTaobaoQimenOrderexceptionReportResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenOrderexceptionReportResponse)
	},
}

// GetTaobaoQimenOrderexceptionReportResponse() 从对象池中获取TaobaoQimenOrderexceptionReportResponse
func GetTaobaoQimenOrderexceptionReportResponse() *TaobaoQimenOrderexceptionReportResponse {
	return poolTaobaoQimenOrderexceptionReportResponse.Get().(*TaobaoQimenOrderexceptionReportResponse)
}

// ReleaseTaobaoQimenOrderexceptionReportResponse 释放TaobaoQimenOrderexceptionReportResponse
func ReleaseTaobaoQimenOrderexceptionReportResponse(v *TaobaoQimenOrderexceptionReportResponse) {
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	poolTaobaoQimenOrderexceptionReportResponse.Put(v)
}
