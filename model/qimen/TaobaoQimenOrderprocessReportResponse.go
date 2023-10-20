package qimen

import (
	"sync"
)

// TaobaoQimenOrderprocessReportResponse 结构体
type TaobaoQimenOrderprocessReportResponse struct {
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolTaobaoQimenOrderprocessReportResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenOrderprocessReportResponse)
	},
}

// GetTaobaoQimenOrderprocessReportResponse() 从对象池中获取TaobaoQimenOrderprocessReportResponse
func GetTaobaoQimenOrderprocessReportResponse() *TaobaoQimenOrderprocessReportResponse {
	return poolTaobaoQimenOrderprocessReportResponse.Get().(*TaobaoQimenOrderprocessReportResponse)
}

// ReleaseTaobaoQimenOrderprocessReportResponse 释放TaobaoQimenOrderprocessReportResponse
func ReleaseTaobaoQimenOrderprocessReportResponse(v *TaobaoQimenOrderprocessReportResponse) {
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	poolTaobaoQimenOrderprocessReportResponse.Put(v)
}
