package simba

import (
	"sync"
)

// TaobaoUniversalbpCampaigngroupFindlistTopResult 结构体
type TaobaoUniversalbpCampaigngroupFindlistTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	CampaignGroupVOTopBulkData *TopBulkData `json:"campaign_group_v_o_top_bulk_data,omitempty" xml:"campaign_group_v_o_top_bulk_data,omitempty"`
}

var poolTaobaoUniversalbpCampaigngroupFindlistTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpCampaigngroupFindlistTopResult)
	},
}

// GetTaobaoUniversalbpCampaigngroupFindlistTopResult() 从对象池中获取TaobaoUniversalbpCampaigngroupFindlistTopResult
func GetTaobaoUniversalbpCampaigngroupFindlistTopResult() *TaobaoUniversalbpCampaigngroupFindlistTopResult {
	return poolTaobaoUniversalbpCampaigngroupFindlistTopResult.Get().(*TaobaoUniversalbpCampaigngroupFindlistTopResult)
}

// ReleaseTaobaoUniversalbpCampaigngroupFindlistTopResult 释放TaobaoUniversalbpCampaigngroupFindlistTopResult
func ReleaseTaobaoUniversalbpCampaigngroupFindlistTopResult(v *TaobaoUniversalbpCampaigngroupFindlistTopResult) {
	v.Info = nil
	v.CampaignGroupVOTopBulkData = nil
	poolTaobaoUniversalbpCampaigngroupFindlistTopResult.Put(v)
}
