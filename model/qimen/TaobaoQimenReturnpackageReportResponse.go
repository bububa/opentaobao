package qimen

import (
	"sync"
)

// TaobaoQimenReturnpackageReportResponse 结构体
type TaobaoQimenReturnpackageReportResponse struct {
	// success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolTaobaoQimenReturnpackageReportResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenReturnpackageReportResponse)
	},
}

// GetTaobaoQimenReturnpackageReportResponse() 从对象池中获取TaobaoQimenReturnpackageReportResponse
func GetTaobaoQimenReturnpackageReportResponse() *TaobaoQimenReturnpackageReportResponse {
	return poolTaobaoQimenReturnpackageReportResponse.Get().(*TaobaoQimenReturnpackageReportResponse)
}

// ReleaseTaobaoQimenReturnpackageReportResponse 释放TaobaoQimenReturnpackageReportResponse
func ReleaseTaobaoQimenReturnpackageReportResponse(v *TaobaoQimenReturnpackageReportResponse) {
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	poolTaobaoQimenReturnpackageReportResponse.Put(v)
}
