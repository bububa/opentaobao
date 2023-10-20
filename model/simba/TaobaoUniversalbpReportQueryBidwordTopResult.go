package simba

import (
	"sync"
)

// TaobaoUniversalbpReportQueryBidwordTopResult 结构体
type TaobaoUniversalbpReportQueryBidwordTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	TopReportVOTopBulkData *TopBulkData `json:"top_report_v_o_top_bulk_data,omitempty" xml:"top_report_v_o_top_bulk_data,omitempty"`
}

var poolTaobaoUniversalbpReportQueryBidwordTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpReportQueryBidwordTopResult)
	},
}

// GetTaobaoUniversalbpReportQueryBidwordTopResult() 从对象池中获取TaobaoUniversalbpReportQueryBidwordTopResult
func GetTaobaoUniversalbpReportQueryBidwordTopResult() *TaobaoUniversalbpReportQueryBidwordTopResult {
	return poolTaobaoUniversalbpReportQueryBidwordTopResult.Get().(*TaobaoUniversalbpReportQueryBidwordTopResult)
}

// ReleaseTaobaoUniversalbpReportQueryBidwordTopResult 释放TaobaoUniversalbpReportQueryBidwordTopResult
func ReleaseTaobaoUniversalbpReportQueryBidwordTopResult(v *TaobaoUniversalbpReportQueryBidwordTopResult) {
	v.Info = nil
	v.TopReportVOTopBulkData = nil
	poolTaobaoUniversalbpReportQueryBidwordTopResult.Put(v)
}
