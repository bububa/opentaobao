package simba

import (
	"sync"
)

// TaobaoUniversalbpStdcategoryFindcategoryconditionTopResult 结构体
type TaobaoUniversalbpStdcategoryFindcategoryconditionTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	StdCategoryVOTopBulkData *TopBulkData `json:"std_category_v_o_top_bulk_data,omitempty" xml:"std_category_v_o_top_bulk_data,omitempty"`
}

var poolTaobaoUniversalbpStdcategoryFindcategoryconditionTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpStdcategoryFindcategoryconditionTopResult)
	},
}

// GetTaobaoUniversalbpStdcategoryFindcategoryconditionTopResult() 从对象池中获取TaobaoUniversalbpStdcategoryFindcategoryconditionTopResult
func GetTaobaoUniversalbpStdcategoryFindcategoryconditionTopResult() *TaobaoUniversalbpStdcategoryFindcategoryconditionTopResult {
	return poolTaobaoUniversalbpStdcategoryFindcategoryconditionTopResult.Get().(*TaobaoUniversalbpStdcategoryFindcategoryconditionTopResult)
}

// ReleaseTaobaoUniversalbpStdcategoryFindcategoryconditionTopResult 释放TaobaoUniversalbpStdcategoryFindcategoryconditionTopResult
func ReleaseTaobaoUniversalbpStdcategoryFindcategoryconditionTopResult(v *TaobaoUniversalbpStdcategoryFindcategoryconditionTopResult) {
	v.Info = nil
	v.StdCategoryVOTopBulkData = nil
	poolTaobaoUniversalbpStdcategoryFindcategoryconditionTopResult.Put(v)
}
