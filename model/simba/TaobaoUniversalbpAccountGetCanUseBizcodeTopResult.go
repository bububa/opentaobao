package simba

import (
	"sync"
)

// TaobaoUniversalbpAccountGetCanUseBizcodeTopResult 结构体
type TaobaoUniversalbpAccountGetCanUseBizcodeTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	TopMarketSceneVOTopBulkData *TopBulkData `json:"top_market_scene_v_o_top_bulk_data,omitempty" xml:"top_market_scene_v_o_top_bulk_data,omitempty"`
}

var poolTaobaoUniversalbpAccountGetCanUseBizcodeTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpAccountGetCanUseBizcodeTopResult)
	},
}

// GetTaobaoUniversalbpAccountGetCanUseBizcodeTopResult() 从对象池中获取TaobaoUniversalbpAccountGetCanUseBizcodeTopResult
func GetTaobaoUniversalbpAccountGetCanUseBizcodeTopResult() *TaobaoUniversalbpAccountGetCanUseBizcodeTopResult {
	return poolTaobaoUniversalbpAccountGetCanUseBizcodeTopResult.Get().(*TaobaoUniversalbpAccountGetCanUseBizcodeTopResult)
}

// ReleaseTaobaoUniversalbpAccountGetCanUseBizcodeTopResult 释放TaobaoUniversalbpAccountGetCanUseBizcodeTopResult
func ReleaseTaobaoUniversalbpAccountGetCanUseBizcodeTopResult(v *TaobaoUniversalbpAccountGetCanUseBizcodeTopResult) {
	v.Info = nil
	v.TopMarketSceneVOTopBulkData = nil
	poolTaobaoUniversalbpAccountGetCanUseBizcodeTopResult.Put(v)
}
