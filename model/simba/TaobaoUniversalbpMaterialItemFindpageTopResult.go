package simba

import (
	"sync"
)

// TaobaoUniversalbpMaterialItemFindpageTopResult 结构体
type TaobaoUniversalbpMaterialItemFindpageTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	ItemVOTopBulkData *TopBulkData `json:"item_v_o_top_bulk_data,omitempty" xml:"item_v_o_top_bulk_data,omitempty"`
}

var poolTaobaoUniversalbpMaterialItemFindpageTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpMaterialItemFindpageTopResult)
	},
}

// GetTaobaoUniversalbpMaterialItemFindpageTopResult() 从对象池中获取TaobaoUniversalbpMaterialItemFindpageTopResult
func GetTaobaoUniversalbpMaterialItemFindpageTopResult() *TaobaoUniversalbpMaterialItemFindpageTopResult {
	return poolTaobaoUniversalbpMaterialItemFindpageTopResult.Get().(*TaobaoUniversalbpMaterialItemFindpageTopResult)
}

// ReleaseTaobaoUniversalbpMaterialItemFindpageTopResult 释放TaobaoUniversalbpMaterialItemFindpageTopResult
func ReleaseTaobaoUniversalbpMaterialItemFindpageTopResult(v *TaobaoUniversalbpMaterialItemFindpageTopResult) {
	v.Info = nil
	v.ItemVOTopBulkData = nil
	poolTaobaoUniversalbpMaterialItemFindpageTopResult.Put(v)
}
