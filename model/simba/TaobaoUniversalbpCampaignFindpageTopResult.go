package simba

import (
	"sync"
)

// TaobaoUniversalbpCampaignFindpageTopResult 结构体
type TaobaoUniversalbpCampaignFindpageTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	CampaignVOTopBulkData *TopBulkData `json:"campaign_v_o_top_bulk_data,omitempty" xml:"campaign_v_o_top_bulk_data,omitempty"`
}

var poolTaobaoUniversalbpCampaignFindpageTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpCampaignFindpageTopResult)
	},
}

// GetTaobaoUniversalbpCampaignFindpageTopResult() 从对象池中获取TaobaoUniversalbpCampaignFindpageTopResult
func GetTaobaoUniversalbpCampaignFindpageTopResult() *TaobaoUniversalbpCampaignFindpageTopResult {
	return poolTaobaoUniversalbpCampaignFindpageTopResult.Get().(*TaobaoUniversalbpCampaignFindpageTopResult)
}

// ReleaseTaobaoUniversalbpCampaignFindpageTopResult 释放TaobaoUniversalbpCampaignFindpageTopResult
func ReleaseTaobaoUniversalbpCampaignFindpageTopResult(v *TaobaoUniversalbpCampaignFindpageTopResult) {
	v.Info = nil
	v.CampaignVOTopBulkData = nil
	poolTaobaoUniversalbpCampaignFindpageTopResult.Put(v)
}
