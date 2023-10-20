package simba

import (
	"sync"
)

// TaobaoUniversalbpReportQueryCrowdTopResult 结构体
type TaobaoUniversalbpReportQueryCrowdTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	TopReportVOTopBulkData *TopBulkData `json:"top_report_v_o_top_bulk_data,omitempty" xml:"top_report_v_o_top_bulk_data,omitempty"`
}

var poolTaobaoUniversalbpReportQueryCrowdTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpReportQueryCrowdTopResult)
	},
}

// GetTaobaoUniversalbpReportQueryCrowdTopResult() 从对象池中获取TaobaoUniversalbpReportQueryCrowdTopResult
func GetTaobaoUniversalbpReportQueryCrowdTopResult() *TaobaoUniversalbpReportQueryCrowdTopResult {
	return poolTaobaoUniversalbpReportQueryCrowdTopResult.Get().(*TaobaoUniversalbpReportQueryCrowdTopResult)
}

// ReleaseTaobaoUniversalbpReportQueryCrowdTopResult 释放TaobaoUniversalbpReportQueryCrowdTopResult
func ReleaseTaobaoUniversalbpReportQueryCrowdTopResult(v *TaobaoUniversalbpReportQueryCrowdTopResult) {
	v.Info = nil
	v.TopReportVOTopBulkData = nil
	poolTaobaoUniversalbpReportQueryCrowdTopResult.Put(v)
}
