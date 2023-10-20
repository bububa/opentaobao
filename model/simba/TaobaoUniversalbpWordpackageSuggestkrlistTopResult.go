package simba

import (
	"sync"
)

// TaobaoUniversalbpWordpackageSuggestkrlistTopResult 结构体
type TaobaoUniversalbpWordpackageSuggestkrlistTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	SuggestWordPackageVOTopBulkData *TopBulkData `json:"suggest_word_package_v_o_top_bulk_data,omitempty" xml:"suggest_word_package_v_o_top_bulk_data,omitempty"`
}

var poolTaobaoUniversalbpWordpackageSuggestkrlistTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpWordpackageSuggestkrlistTopResult)
	},
}

// GetTaobaoUniversalbpWordpackageSuggestkrlistTopResult() 从对象池中获取TaobaoUniversalbpWordpackageSuggestkrlistTopResult
func GetTaobaoUniversalbpWordpackageSuggestkrlistTopResult() *TaobaoUniversalbpWordpackageSuggestkrlistTopResult {
	return poolTaobaoUniversalbpWordpackageSuggestkrlistTopResult.Get().(*TaobaoUniversalbpWordpackageSuggestkrlistTopResult)
}

// ReleaseTaobaoUniversalbpWordpackageSuggestkrlistTopResult 释放TaobaoUniversalbpWordpackageSuggestkrlistTopResult
func ReleaseTaobaoUniversalbpWordpackageSuggestkrlistTopResult(v *TaobaoUniversalbpWordpackageSuggestkrlistTopResult) {
	v.Info = nil
	v.SuggestWordPackageVOTopBulkData = nil
	poolTaobaoUniversalbpWordpackageSuggestkrlistTopResult.Put(v)
}
