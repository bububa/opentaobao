package simba

import (
	"sync"
)

// TaobaoUniversalbpBidwordSuggestdefaultlistTopResult 结构体
type TaobaoUniversalbpBidwordSuggestdefaultlistTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	BidwordSuggestItemVOTopBulkData *TopBulkData `json:"bidword_suggest_item_v_o_top_bulk_data,omitempty" xml:"bidword_suggest_item_v_o_top_bulk_data,omitempty"`
}

var poolTaobaoUniversalbpBidwordSuggestdefaultlistTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpBidwordSuggestdefaultlistTopResult)
	},
}

// GetTaobaoUniversalbpBidwordSuggestdefaultlistTopResult() 从对象池中获取TaobaoUniversalbpBidwordSuggestdefaultlistTopResult
func GetTaobaoUniversalbpBidwordSuggestdefaultlistTopResult() *TaobaoUniversalbpBidwordSuggestdefaultlistTopResult {
	return poolTaobaoUniversalbpBidwordSuggestdefaultlistTopResult.Get().(*TaobaoUniversalbpBidwordSuggestdefaultlistTopResult)
}

// ReleaseTaobaoUniversalbpBidwordSuggestdefaultlistTopResult 释放TaobaoUniversalbpBidwordSuggestdefaultlistTopResult
func ReleaseTaobaoUniversalbpBidwordSuggestdefaultlistTopResult(v *TaobaoUniversalbpBidwordSuggestdefaultlistTopResult) {
	v.Info = nil
	v.BidwordSuggestItemVOTopBulkData = nil
	poolTaobaoUniversalbpBidwordSuggestdefaultlistTopResult.Put(v)
}
