package simba

import (
	"sync"
)

// TaobaoUniversalbpWordpackageSuggestdefaultlistTopResult 结构体
type TaobaoUniversalbpWordpackageSuggestdefaultlistTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	WordPackageSuggestItemVOTopBulkData *TopBulkData `json:"word_package_suggest_item_v_o_top_bulk_data,omitempty" xml:"word_package_suggest_item_v_o_top_bulk_data,omitempty"`
}

var poolTaobaoUniversalbpWordpackageSuggestdefaultlistTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpWordpackageSuggestdefaultlistTopResult)
	},
}

// GetTaobaoUniversalbpWordpackageSuggestdefaultlistTopResult() 从对象池中获取TaobaoUniversalbpWordpackageSuggestdefaultlistTopResult
func GetTaobaoUniversalbpWordpackageSuggestdefaultlistTopResult() *TaobaoUniversalbpWordpackageSuggestdefaultlistTopResult {
	return poolTaobaoUniversalbpWordpackageSuggestdefaultlistTopResult.Get().(*TaobaoUniversalbpWordpackageSuggestdefaultlistTopResult)
}

// ReleaseTaobaoUniversalbpWordpackageSuggestdefaultlistTopResult 释放TaobaoUniversalbpWordpackageSuggestdefaultlistTopResult
func ReleaseTaobaoUniversalbpWordpackageSuggestdefaultlistTopResult(v *TaobaoUniversalbpWordpackageSuggestdefaultlistTopResult) {
	v.Info = nil
	v.WordPackageSuggestItemVOTopBulkData = nil
	poolTaobaoUniversalbpWordpackageSuggestdefaultlistTopResult.Put(v)
}
