package simba

import (
	"sync"
)

// TaobaoUniversalbpReportQueryCampaignTopResult 结构体
type TaobaoUniversalbpReportQueryCampaignTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	TopReportVOTopBulkData *TopBulkData `json:"top_report_v_o_top_bulk_data,omitempty" xml:"top_report_v_o_top_bulk_data,omitempty"`
}

var poolTaobaoUniversalbpReportQueryCampaignTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpReportQueryCampaignTopResult)
	},
}

// GetTaobaoUniversalbpReportQueryCampaignTopResult() 从对象池中获取TaobaoUniversalbpReportQueryCampaignTopResult
func GetTaobaoUniversalbpReportQueryCampaignTopResult() *TaobaoUniversalbpReportQueryCampaignTopResult {
	return poolTaobaoUniversalbpReportQueryCampaignTopResult.Get().(*TaobaoUniversalbpReportQueryCampaignTopResult)
}

// ReleaseTaobaoUniversalbpReportQueryCampaignTopResult 释放TaobaoUniversalbpReportQueryCampaignTopResult
func ReleaseTaobaoUniversalbpReportQueryCampaignTopResult(v *TaobaoUniversalbpReportQueryCampaignTopResult) {
	v.Info = nil
	v.TopReportVOTopBulkData = nil
	poolTaobaoUniversalbpReportQueryCampaignTopResult.Put(v)
}
