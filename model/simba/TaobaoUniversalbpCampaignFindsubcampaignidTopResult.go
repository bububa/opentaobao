package simba

import (
	"sync"
)

// TaobaoUniversalbpCampaignFindsubcampaignidTopResult 结构体
type TaobaoUniversalbpCampaignFindsubcampaignidTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
}

var poolTaobaoUniversalbpCampaignFindsubcampaignidTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpCampaignFindsubcampaignidTopResult)
	},
}

// GetTaobaoUniversalbpCampaignFindsubcampaignidTopResult() 从对象池中获取TaobaoUniversalbpCampaignFindsubcampaignidTopResult
func GetTaobaoUniversalbpCampaignFindsubcampaignidTopResult() *TaobaoUniversalbpCampaignFindsubcampaignidTopResult {
	return poolTaobaoUniversalbpCampaignFindsubcampaignidTopResult.Get().(*TaobaoUniversalbpCampaignFindsubcampaignidTopResult)
}

// ReleaseTaobaoUniversalbpCampaignFindsubcampaignidTopResult 释放TaobaoUniversalbpCampaignFindsubcampaignidTopResult
func ReleaseTaobaoUniversalbpCampaignFindsubcampaignidTopResult(v *TaobaoUniversalbpCampaignFindsubcampaignidTopResult) {
	v.Info = nil
	v.Data = 0
	poolTaobaoUniversalbpCampaignFindsubcampaignidTopResult.Put(v)
}
