package simba

import (
	"sync"
)

// TaobaoUniversalbpWordpackageFindlistTopResult 结构体
type TaobaoUniversalbpWordpackageFindlistTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	WordPackageVOTopBulkData *TopBulkData `json:"word_package_v_o_top_bulk_data,omitempty" xml:"word_package_v_o_top_bulk_data,omitempty"`
}

var poolTaobaoUniversalbpWordpackageFindlistTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpWordpackageFindlistTopResult)
	},
}

// GetTaobaoUniversalbpWordpackageFindlistTopResult() 从对象池中获取TaobaoUniversalbpWordpackageFindlistTopResult
func GetTaobaoUniversalbpWordpackageFindlistTopResult() *TaobaoUniversalbpWordpackageFindlistTopResult {
	return poolTaobaoUniversalbpWordpackageFindlistTopResult.Get().(*TaobaoUniversalbpWordpackageFindlistTopResult)
}

// ReleaseTaobaoUniversalbpWordpackageFindlistTopResult 释放TaobaoUniversalbpWordpackageFindlistTopResult
func ReleaseTaobaoUniversalbpWordpackageFindlistTopResult(v *TaobaoUniversalbpWordpackageFindlistTopResult) {
	v.Info = nil
	v.WordPackageVOTopBulkData = nil
	poolTaobaoUniversalbpWordpackageFindlistTopResult.Put(v)
}
