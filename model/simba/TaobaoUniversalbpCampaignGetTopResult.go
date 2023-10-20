package simba

import (
	"sync"
)

// TaobaoUniversalbpCampaignGetTopResult 结构体
type TaobaoUniversalbpCampaignGetTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	TopCampaignVO *TopCampaignVo `json:"top_campaign_v_o,omitempty" xml:"top_campaign_v_o,omitempty"`
}

var poolTaobaoUniversalbpCampaignGetTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpCampaignGetTopResult)
	},
}

// GetTaobaoUniversalbpCampaignGetTopResult() 从对象池中获取TaobaoUniversalbpCampaignGetTopResult
func GetTaobaoUniversalbpCampaignGetTopResult() *TaobaoUniversalbpCampaignGetTopResult {
	return poolTaobaoUniversalbpCampaignGetTopResult.Get().(*TaobaoUniversalbpCampaignGetTopResult)
}

// ReleaseTaobaoUniversalbpCampaignGetTopResult 释放TaobaoUniversalbpCampaignGetTopResult
func ReleaseTaobaoUniversalbpCampaignGetTopResult(v *TaobaoUniversalbpCampaignGetTopResult) {
	v.Info = nil
	v.TopCampaignVO = nil
	poolTaobaoUniversalbpCampaignGetTopResult.Put(v)
}
