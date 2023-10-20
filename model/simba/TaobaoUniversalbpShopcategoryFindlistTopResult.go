package simba

import (
	"sync"
)

// TaobaoUniversalbpShopcategoryFindlistTopResult 结构体
type TaobaoUniversalbpShopcategoryFindlistTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	ShopCategoryVOTopBulkData *TopBulkData `json:"shop_category_v_o_top_bulk_data,omitempty" xml:"shop_category_v_o_top_bulk_data,omitempty"`
}

var poolTaobaoUniversalbpShopcategoryFindlistTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpShopcategoryFindlistTopResult)
	},
}

// GetTaobaoUniversalbpShopcategoryFindlistTopResult() 从对象池中获取TaobaoUniversalbpShopcategoryFindlistTopResult
func GetTaobaoUniversalbpShopcategoryFindlistTopResult() *TaobaoUniversalbpShopcategoryFindlistTopResult {
	return poolTaobaoUniversalbpShopcategoryFindlistTopResult.Get().(*TaobaoUniversalbpShopcategoryFindlistTopResult)
}

// ReleaseTaobaoUniversalbpShopcategoryFindlistTopResult 释放TaobaoUniversalbpShopcategoryFindlistTopResult
func ReleaseTaobaoUniversalbpShopcategoryFindlistTopResult(v *TaobaoUniversalbpShopcategoryFindlistTopResult) {
	v.Info = nil
	v.ShopCategoryVOTopBulkData = nil
	poolTaobaoUniversalbpShopcategoryFindlistTopResult.Put(v)
}
