package simba

import (
	"sync"
)

// TaobaoUniversalbpReportQueryNotItemPromotionTopResult 结构体
type TaobaoUniversalbpReportQueryNotItemPromotionTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	TopReportVOTopBulkData *TopBulkData `json:"top_report_v_o_top_bulk_data,omitempty" xml:"top_report_v_o_top_bulk_data,omitempty"`
}

var poolTaobaoUniversalbpReportQueryNotItemPromotionTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpReportQueryNotItemPromotionTopResult)
	},
}

// GetTaobaoUniversalbpReportQueryNotItemPromotionTopResult() 从对象池中获取TaobaoUniversalbpReportQueryNotItemPromotionTopResult
func GetTaobaoUniversalbpReportQueryNotItemPromotionTopResult() *TaobaoUniversalbpReportQueryNotItemPromotionTopResult {
	return poolTaobaoUniversalbpReportQueryNotItemPromotionTopResult.Get().(*TaobaoUniversalbpReportQueryNotItemPromotionTopResult)
}

// ReleaseTaobaoUniversalbpReportQueryNotItemPromotionTopResult 释放TaobaoUniversalbpReportQueryNotItemPromotionTopResult
func ReleaseTaobaoUniversalbpReportQueryNotItemPromotionTopResult(v *TaobaoUniversalbpReportQueryNotItemPromotionTopResult) {
	v.Info = nil
	v.TopReportVOTopBulkData = nil
	poolTaobaoUniversalbpReportQueryNotItemPromotionTopResult.Put(v)
}
