package qimen

import (
	"sync"
)

// TaobaoQimenStockchangeReportResponse 结构体
type TaobaoQimenStockchangeReportResponse struct {
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolTaobaoQimenStockchangeReportResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenStockchangeReportResponse)
	},
}

// GetTaobaoQimenStockchangeReportResponse() 从对象池中获取TaobaoQimenStockchangeReportResponse
func GetTaobaoQimenStockchangeReportResponse() *TaobaoQimenStockchangeReportResponse {
	return poolTaobaoQimenStockchangeReportResponse.Get().(*TaobaoQimenStockchangeReportResponse)
}

// ReleaseTaobaoQimenStockchangeReportResponse 释放TaobaoQimenStockchangeReportResponse
func ReleaseTaobaoQimenStockchangeReportResponse(v *TaobaoQimenStockchangeReportResponse) {
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	poolTaobaoQimenStockchangeReportResponse.Put(v)
}
