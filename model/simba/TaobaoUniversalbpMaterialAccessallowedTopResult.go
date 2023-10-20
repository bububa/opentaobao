package simba

import (
	"sync"
)

// TaobaoUniversalbpMaterialAccessallowedTopResult 结构体
type TaobaoUniversalbpMaterialAccessallowedTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	MaterialAccessAllowVOTopBulkData *TopBulkData `json:"material_access_allow_v_o_top_bulk_data,omitempty" xml:"material_access_allow_v_o_top_bulk_data,omitempty"`
}

var poolTaobaoUniversalbpMaterialAccessallowedTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpMaterialAccessallowedTopResult)
	},
}

// GetTaobaoUniversalbpMaterialAccessallowedTopResult() 从对象池中获取TaobaoUniversalbpMaterialAccessallowedTopResult
func GetTaobaoUniversalbpMaterialAccessallowedTopResult() *TaobaoUniversalbpMaterialAccessallowedTopResult {
	return poolTaobaoUniversalbpMaterialAccessallowedTopResult.Get().(*TaobaoUniversalbpMaterialAccessallowedTopResult)
}

// ReleaseTaobaoUniversalbpMaterialAccessallowedTopResult 释放TaobaoUniversalbpMaterialAccessallowedTopResult
func ReleaseTaobaoUniversalbpMaterialAccessallowedTopResult(v *TaobaoUniversalbpMaterialAccessallowedTopResult) {
	v.Info = nil
	v.MaterialAccessAllowVOTopBulkData = nil
	poolTaobaoUniversalbpMaterialAccessallowedTopResult.Put(v)
}
