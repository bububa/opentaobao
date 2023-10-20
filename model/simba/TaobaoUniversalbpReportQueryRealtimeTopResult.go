package simba

import (
	"sync"
)

// TaobaoUniversalbpReportQueryRealtimeTopResult 结构体
type TaobaoUniversalbpReportQueryRealtimeTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	TopReportVOTopBulkData *TopBulkData `json:"top_report_v_o_top_bulk_data,omitempty" xml:"top_report_v_o_top_bulk_data,omitempty"`
}

var poolTaobaoUniversalbpReportQueryRealtimeTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpReportQueryRealtimeTopResult)
	},
}

// GetTaobaoUniversalbpReportQueryRealtimeTopResult() 从对象池中获取TaobaoUniversalbpReportQueryRealtimeTopResult
func GetTaobaoUniversalbpReportQueryRealtimeTopResult() *TaobaoUniversalbpReportQueryRealtimeTopResult {
	return poolTaobaoUniversalbpReportQueryRealtimeTopResult.Get().(*TaobaoUniversalbpReportQueryRealtimeTopResult)
}

// ReleaseTaobaoUniversalbpReportQueryRealtimeTopResult 释放TaobaoUniversalbpReportQueryRealtimeTopResult
func ReleaseTaobaoUniversalbpReportQueryRealtimeTopResult(v *TaobaoUniversalbpReportQueryRealtimeTopResult) {
	v.Info = nil
	v.TopReportVOTopBulkData = nil
	poolTaobaoUniversalbpReportQueryRealtimeTopResult.Put(v)
}
