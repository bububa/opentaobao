package simba

import (
	"sync"
)

// TaobaoUniversalbpReportQueryAdgroupTopResult 结构体
type TaobaoUniversalbpReportQueryAdgroupTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	TopReportVOTopBulkData *TopBulkData `json:"top_report_v_o_top_bulk_data,omitempty" xml:"top_report_v_o_top_bulk_data,omitempty"`
}

var poolTaobaoUniversalbpReportQueryAdgroupTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpReportQueryAdgroupTopResult)
	},
}

// GetTaobaoUniversalbpReportQueryAdgroupTopResult() 从对象池中获取TaobaoUniversalbpReportQueryAdgroupTopResult
func GetTaobaoUniversalbpReportQueryAdgroupTopResult() *TaobaoUniversalbpReportQueryAdgroupTopResult {
	return poolTaobaoUniversalbpReportQueryAdgroupTopResult.Get().(*TaobaoUniversalbpReportQueryAdgroupTopResult)
}

// ReleaseTaobaoUniversalbpReportQueryAdgroupTopResult 释放TaobaoUniversalbpReportQueryAdgroupTopResult
func ReleaseTaobaoUniversalbpReportQueryAdgroupTopResult(v *TaobaoUniversalbpReportQueryAdgroupTopResult) {
	v.Info = nil
	v.TopReportVOTopBulkData = nil
	poolTaobaoUniversalbpReportQueryAdgroupTopResult.Put(v)
}
