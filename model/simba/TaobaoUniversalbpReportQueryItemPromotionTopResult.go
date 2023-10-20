package simba

import (
	"sync"
)

// TaobaoUniversalbpReportQueryItemPromotionTopResult 结构体
type TaobaoUniversalbpReportQueryItemPromotionTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	TopReportVOTopBulkData *TopBulkData `json:"top_report_v_o_top_bulk_data,omitempty" xml:"top_report_v_o_top_bulk_data,omitempty"`
}

var poolTaobaoUniversalbpReportQueryItemPromotionTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpReportQueryItemPromotionTopResult)
	},
}

// GetTaobaoUniversalbpReportQueryItemPromotionTopResult() 从对象池中获取TaobaoUniversalbpReportQueryItemPromotionTopResult
func GetTaobaoUniversalbpReportQueryItemPromotionTopResult() *TaobaoUniversalbpReportQueryItemPromotionTopResult {
	return poolTaobaoUniversalbpReportQueryItemPromotionTopResult.Get().(*TaobaoUniversalbpReportQueryItemPromotionTopResult)
}

// ReleaseTaobaoUniversalbpReportQueryItemPromotionTopResult 释放TaobaoUniversalbpReportQueryItemPromotionTopResult
func ReleaseTaobaoUniversalbpReportQueryItemPromotionTopResult(v *TaobaoUniversalbpReportQueryItemPromotionTopResult) {
	v.Info = nil
	v.TopReportVOTopBulkData = nil
	poolTaobaoUniversalbpReportQueryItemPromotionTopResult.Put(v)
}
