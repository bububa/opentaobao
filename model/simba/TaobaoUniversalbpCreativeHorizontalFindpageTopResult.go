package simba

import (
	"sync"
)

// TaobaoUniversalbpCreativeHorizontalFindpageTopResult 结构体
type TaobaoUniversalbpCreativeHorizontalFindpageTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	CreativeRefVOTopBulkData *TopBulkData `json:"creative_ref_v_o_top_bulk_data,omitempty" xml:"creative_ref_v_o_top_bulk_data,omitempty"`
}

var poolTaobaoUniversalbpCreativeHorizontalFindpageTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpCreativeHorizontalFindpageTopResult)
	},
}

// GetTaobaoUniversalbpCreativeHorizontalFindpageTopResult() 从对象池中获取TaobaoUniversalbpCreativeHorizontalFindpageTopResult
func GetTaobaoUniversalbpCreativeHorizontalFindpageTopResult() *TaobaoUniversalbpCreativeHorizontalFindpageTopResult {
	return poolTaobaoUniversalbpCreativeHorizontalFindpageTopResult.Get().(*TaobaoUniversalbpCreativeHorizontalFindpageTopResult)
}

// ReleaseTaobaoUniversalbpCreativeHorizontalFindpageTopResult 释放TaobaoUniversalbpCreativeHorizontalFindpageTopResult
func ReleaseTaobaoUniversalbpCreativeHorizontalFindpageTopResult(v *TaobaoUniversalbpCreativeHorizontalFindpageTopResult) {
	v.Info = nil
	v.CreativeRefVOTopBulkData = nil
	poolTaobaoUniversalbpCreativeHorizontalFindpageTopResult.Put(v)
}
