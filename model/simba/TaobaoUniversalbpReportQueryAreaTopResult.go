package simba

import (
	"sync"
)

// TaobaoUniversalbpReportQueryAreaTopResult 结构体
type TaobaoUniversalbpReportQueryAreaTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	TopReportVOTopBulkData *TopBulkData `json:"top_report_v_o_top_bulk_data,omitempty" xml:"top_report_v_o_top_bulk_data,omitempty"`
}

var poolTaobaoUniversalbpReportQueryAreaTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpReportQueryAreaTopResult)
	},
}

// GetTaobaoUniversalbpReportQueryAreaTopResult() 从对象池中获取TaobaoUniversalbpReportQueryAreaTopResult
func GetTaobaoUniversalbpReportQueryAreaTopResult() *TaobaoUniversalbpReportQueryAreaTopResult {
	return poolTaobaoUniversalbpReportQueryAreaTopResult.Get().(*TaobaoUniversalbpReportQueryAreaTopResult)
}

// ReleaseTaobaoUniversalbpReportQueryAreaTopResult 释放TaobaoUniversalbpReportQueryAreaTopResult
func ReleaseTaobaoUniversalbpReportQueryAreaTopResult(v *TaobaoUniversalbpReportQueryAreaTopResult) {
	v.Info = nil
	v.TopReportVOTopBulkData = nil
	poolTaobaoUniversalbpReportQueryAreaTopResult.Put(v)
}
