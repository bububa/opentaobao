package simba

import (
	"sync"
)

// TaobaoUniversalbpStdcategoryFindlistTopResult 结构体
type TaobaoUniversalbpStdcategoryFindlistTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	StdCategoryVOTopBulkData *TopBulkData `json:"std_category_v_o_top_bulk_data,omitempty" xml:"std_category_v_o_top_bulk_data,omitempty"`
}

var poolTaobaoUniversalbpStdcategoryFindlistTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpStdcategoryFindlistTopResult)
	},
}

// GetTaobaoUniversalbpStdcategoryFindlistTopResult() 从对象池中获取TaobaoUniversalbpStdcategoryFindlistTopResult
func GetTaobaoUniversalbpStdcategoryFindlistTopResult() *TaobaoUniversalbpStdcategoryFindlistTopResult {
	return poolTaobaoUniversalbpStdcategoryFindlistTopResult.Get().(*TaobaoUniversalbpStdcategoryFindlistTopResult)
}

// ReleaseTaobaoUniversalbpStdcategoryFindlistTopResult 释放TaobaoUniversalbpStdcategoryFindlistTopResult
func ReleaseTaobaoUniversalbpStdcategoryFindlistTopResult(v *TaobaoUniversalbpStdcategoryFindlistTopResult) {
	v.Info = nil
	v.StdCategoryVOTopBulkData = nil
	poolTaobaoUniversalbpStdcategoryFindlistTopResult.Put(v)
}
