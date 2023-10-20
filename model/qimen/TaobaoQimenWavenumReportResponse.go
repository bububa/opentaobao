package qimen

import (
	"sync"
)

// TaobaoQimenWavenumReportResponse 结构体
type TaobaoQimenWavenumReportResponse struct {
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolTaobaoQimenWavenumReportResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenWavenumReportResponse)
	},
}

// GetTaobaoQimenWavenumReportResponse() 从对象池中获取TaobaoQimenWavenumReportResponse
func GetTaobaoQimenWavenumReportResponse() *TaobaoQimenWavenumReportResponse {
	return poolTaobaoQimenWavenumReportResponse.Get().(*TaobaoQimenWavenumReportResponse)
}

// ReleaseTaobaoQimenWavenumReportResponse 释放TaobaoQimenWavenumReportResponse
func ReleaseTaobaoQimenWavenumReportResponse(v *TaobaoQimenWavenumReportResponse) {
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	poolTaobaoQimenWavenumReportResponse.Put(v)
}
