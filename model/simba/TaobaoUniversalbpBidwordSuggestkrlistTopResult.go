package simba

import (
	"sync"
)

// TaobaoUniversalbpBidwordSuggestkrlistTopResult 结构体
type TaobaoUniversalbpBidwordSuggestkrlistTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	SuggestBidwordVOTopBulkData *TopBulkData `json:"suggest_bidword_v_o_top_bulk_data,omitempty" xml:"suggest_bidword_v_o_top_bulk_data,omitempty"`
}

var poolTaobaoUniversalbpBidwordSuggestkrlistTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpBidwordSuggestkrlistTopResult)
	},
}

// GetTaobaoUniversalbpBidwordSuggestkrlistTopResult() 从对象池中获取TaobaoUniversalbpBidwordSuggestkrlistTopResult
func GetTaobaoUniversalbpBidwordSuggestkrlistTopResult() *TaobaoUniversalbpBidwordSuggestkrlistTopResult {
	return poolTaobaoUniversalbpBidwordSuggestkrlistTopResult.Get().(*TaobaoUniversalbpBidwordSuggestkrlistTopResult)
}

// ReleaseTaobaoUniversalbpBidwordSuggestkrlistTopResult 释放TaobaoUniversalbpBidwordSuggestkrlistTopResult
func ReleaseTaobaoUniversalbpBidwordSuggestkrlistTopResult(v *TaobaoUniversalbpBidwordSuggestkrlistTopResult) {
	v.Info = nil
	v.SuggestBidwordVOTopBulkData = nil
	poolTaobaoUniversalbpBidwordSuggestkrlistTopResult.Put(v)
}
