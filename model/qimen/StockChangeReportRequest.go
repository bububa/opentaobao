package qimen

import (
	"sync"
)

// StockChangeReportRequest 结构体
type StockChangeReportRequest struct {
	// item
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
	// 扩展属性
	SnList []SnList `json:"snList,omitempty" xml:"snList>sn_list,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenStockchangeReportMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}

var poolStockChangeReportRequest = sync.Pool{
	New: func() any {
		return new(StockChangeReportRequest)
	},
}

// GetStockChangeReportRequest() 从对象池中获取StockChangeReportRequest
func GetStockChangeReportRequest() *StockChangeReportRequest {
	return poolStockChangeReportRequest.Get().(*StockChangeReportRequest)
}

// ReleaseStockChangeReportRequest 释放StockChangeReportRequest
func ReleaseStockChangeReportRequest(v *StockChangeReportRequest) {
	v.Items = v.Items[:0]
	v.SnList = v.SnList[:0]
	v.ExtendProps = nil
	poolStockChangeReportRequest.Put(v)
}
