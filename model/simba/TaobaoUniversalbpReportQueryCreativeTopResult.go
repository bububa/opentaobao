package simba

import (
	"sync"
)

// TaobaoUniversalbpReportQueryCreativeTopResult 结构体
type TaobaoUniversalbpReportQueryCreativeTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	TopReportVOTopBulkData *TopBulkData `json:"top_report_v_o_top_bulk_data,omitempty" xml:"top_report_v_o_top_bulk_data,omitempty"`
}

var poolTaobaoUniversalbpReportQueryCreativeTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpReportQueryCreativeTopResult)
	},
}

// GetTaobaoUniversalbpReportQueryCreativeTopResult() 从对象池中获取TaobaoUniversalbpReportQueryCreativeTopResult
func GetTaobaoUniversalbpReportQueryCreativeTopResult() *TaobaoUniversalbpReportQueryCreativeTopResult {
	return poolTaobaoUniversalbpReportQueryCreativeTopResult.Get().(*TaobaoUniversalbpReportQueryCreativeTopResult)
}

// ReleaseTaobaoUniversalbpReportQueryCreativeTopResult 释放TaobaoUniversalbpReportQueryCreativeTopResult
func ReleaseTaobaoUniversalbpReportQueryCreativeTopResult(v *TaobaoUniversalbpReportQueryCreativeTopResult) {
	v.Info = nil
	v.TopReportVOTopBulkData = nil
	poolTaobaoUniversalbpReportQueryCreativeTopResult.Put(v)
}
