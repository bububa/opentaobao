package simba

import (
	"sync"
)

// TaobaoUniversalbpBidwordFindlistTopResult 结构体
type TaobaoUniversalbpBidwordFindlistTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	WordVOTopBulkData *TopBulkData `json:"word_v_o_top_bulk_data,omitempty" xml:"word_v_o_top_bulk_data,omitempty"`
}

var poolTaobaoUniversalbpBidwordFindlistTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpBidwordFindlistTopResult)
	},
}

// GetTaobaoUniversalbpBidwordFindlistTopResult() 从对象池中获取TaobaoUniversalbpBidwordFindlistTopResult
func GetTaobaoUniversalbpBidwordFindlistTopResult() *TaobaoUniversalbpBidwordFindlistTopResult {
	return poolTaobaoUniversalbpBidwordFindlistTopResult.Get().(*TaobaoUniversalbpBidwordFindlistTopResult)
}

// ReleaseTaobaoUniversalbpBidwordFindlistTopResult 释放TaobaoUniversalbpBidwordFindlistTopResult
func ReleaseTaobaoUniversalbpBidwordFindlistTopResult(v *TaobaoUniversalbpBidwordFindlistTopResult) {
	v.Info = nil
	v.WordVOTopBulkData = nil
	poolTaobaoUniversalbpBidwordFindlistTopResult.Put(v)
}
