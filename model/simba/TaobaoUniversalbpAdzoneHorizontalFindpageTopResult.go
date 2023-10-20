package simba

import (
	"sync"
)

// TaobaoUniversalbpAdzoneHorizontalFindpageTopResult 结构体
type TaobaoUniversalbpAdzoneHorizontalFindpageTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	AdzoneRefVOTopBulkData *TopBulkData `json:"adzone_ref_v_o_top_bulk_data,omitempty" xml:"adzone_ref_v_o_top_bulk_data,omitempty"`
}

var poolTaobaoUniversalbpAdzoneHorizontalFindpageTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpAdzoneHorizontalFindpageTopResult)
	},
}

// GetTaobaoUniversalbpAdzoneHorizontalFindpageTopResult() 从对象池中获取TaobaoUniversalbpAdzoneHorizontalFindpageTopResult
func GetTaobaoUniversalbpAdzoneHorizontalFindpageTopResult() *TaobaoUniversalbpAdzoneHorizontalFindpageTopResult {
	return poolTaobaoUniversalbpAdzoneHorizontalFindpageTopResult.Get().(*TaobaoUniversalbpAdzoneHorizontalFindpageTopResult)
}

// ReleaseTaobaoUniversalbpAdzoneHorizontalFindpageTopResult 释放TaobaoUniversalbpAdzoneHorizontalFindpageTopResult
func ReleaseTaobaoUniversalbpAdzoneHorizontalFindpageTopResult(v *TaobaoUniversalbpAdzoneHorizontalFindpageTopResult) {
	v.Info = nil
	v.AdzoneRefVOTopBulkData = nil
	poolTaobaoUniversalbpAdzoneHorizontalFindpageTopResult.Put(v)
}
