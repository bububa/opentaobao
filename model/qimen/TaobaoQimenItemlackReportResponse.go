package qimen

import (
	"sync"
)

// TaobaoQimenItemlackReportResponse 结构体
type TaobaoQimenItemlackReportResponse struct {
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolTaobaoQimenItemlackReportResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenItemlackReportResponse)
	},
}

// GetTaobaoQimenItemlackReportResponse() 从对象池中获取TaobaoQimenItemlackReportResponse
func GetTaobaoQimenItemlackReportResponse() *TaobaoQimenItemlackReportResponse {
	return poolTaobaoQimenItemlackReportResponse.Get().(*TaobaoQimenItemlackReportResponse)
}

// ReleaseTaobaoQimenItemlackReportResponse 释放TaobaoQimenItemlackReportResponse
func ReleaseTaobaoQimenItemlackReportResponse(v *TaobaoQimenItemlackReportResponse) {
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	poolTaobaoQimenItemlackReportResponse.Put(v)
}
