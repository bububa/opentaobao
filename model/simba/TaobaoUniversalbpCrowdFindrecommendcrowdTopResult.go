package simba

import (
	"sync"
)

// TaobaoUniversalbpCrowdFindrecommendcrowdTopResult 结构体
type TaobaoUniversalbpCrowdFindrecommendcrowdTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	CrowdRefVOTopBulkData *TopBulkData `json:"crowd_ref_v_o_top_bulk_data,omitempty" xml:"crowd_ref_v_o_top_bulk_data,omitempty"`
}

var poolTaobaoUniversalbpCrowdFindrecommendcrowdTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpCrowdFindrecommendcrowdTopResult)
	},
}

// GetTaobaoUniversalbpCrowdFindrecommendcrowdTopResult() 从对象池中获取TaobaoUniversalbpCrowdFindrecommendcrowdTopResult
func GetTaobaoUniversalbpCrowdFindrecommendcrowdTopResult() *TaobaoUniversalbpCrowdFindrecommendcrowdTopResult {
	return poolTaobaoUniversalbpCrowdFindrecommendcrowdTopResult.Get().(*TaobaoUniversalbpCrowdFindrecommendcrowdTopResult)
}

// ReleaseTaobaoUniversalbpCrowdFindrecommendcrowdTopResult 释放TaobaoUniversalbpCrowdFindrecommendcrowdTopResult
func ReleaseTaobaoUniversalbpCrowdFindrecommendcrowdTopResult(v *TaobaoUniversalbpCrowdFindrecommendcrowdTopResult) {
	v.Info = nil
	v.CrowdRefVOTopBulkData = nil
	poolTaobaoUniversalbpCrowdFindrecommendcrowdTopResult.Put(v)
}
