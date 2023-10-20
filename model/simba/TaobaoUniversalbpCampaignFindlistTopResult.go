package simba

import (
	"sync"
)

// TaobaoUniversalbpCampaignFindlistTopResult 结构体
type TaobaoUniversalbpCampaignFindlistTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	CampaignVOTopBulkData *TopBulkData `json:"campaign_v_o_top_bulk_data,omitempty" xml:"campaign_v_o_top_bulk_data,omitempty"`
}

var poolTaobaoUniversalbpCampaignFindlistTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpCampaignFindlistTopResult)
	},
}

// GetTaobaoUniversalbpCampaignFindlistTopResult() 从对象池中获取TaobaoUniversalbpCampaignFindlistTopResult
func GetTaobaoUniversalbpCampaignFindlistTopResult() *TaobaoUniversalbpCampaignFindlistTopResult {
	return poolTaobaoUniversalbpCampaignFindlistTopResult.Get().(*TaobaoUniversalbpCampaignFindlistTopResult)
}

// ReleaseTaobaoUniversalbpCampaignFindlistTopResult 释放TaobaoUniversalbpCampaignFindlistTopResult
func ReleaseTaobaoUniversalbpCampaignFindlistTopResult(v *TaobaoUniversalbpCampaignFindlistTopResult) {
	v.Info = nil
	v.CampaignVOTopBulkData = nil
	poolTaobaoUniversalbpCampaignFindlistTopResult.Put(v)
}
