package simba

import (
	"sync"
)

// TaobaoUniversalbpReportQueryAccountTopResult 结构体
type TaobaoUniversalbpReportQueryAccountTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	TopReportVOTopBulkData *TopBulkData `json:"top_report_v_o_top_bulk_data,omitempty" xml:"top_report_v_o_top_bulk_data,omitempty"`
}

var poolTaobaoUniversalbpReportQueryAccountTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpReportQueryAccountTopResult)
	},
}

// GetTaobaoUniversalbpReportQueryAccountTopResult() 从对象池中获取TaobaoUniversalbpReportQueryAccountTopResult
func GetTaobaoUniversalbpReportQueryAccountTopResult() *TaobaoUniversalbpReportQueryAccountTopResult {
	return poolTaobaoUniversalbpReportQueryAccountTopResult.Get().(*TaobaoUniversalbpReportQueryAccountTopResult)
}

// ReleaseTaobaoUniversalbpReportQueryAccountTopResult 释放TaobaoUniversalbpReportQueryAccountTopResult
func ReleaseTaobaoUniversalbpReportQueryAccountTopResult(v *TaobaoUniversalbpReportQueryAccountTopResult) {
	v.Info = nil
	v.TopReportVOTopBulkData = nil
	poolTaobaoUniversalbpReportQueryAccountTopResult.Put(v)
}
