package simba

import (
	"sync"
)

// TaobaoUniversalbpCrowdFindlistTopResult 结构体
type TaobaoUniversalbpCrowdFindlistTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	CrowdBindResultVOTopBulkData *TopBulkData `json:"crowd_bind_result_v_o_top_bulk_data,omitempty" xml:"crowd_bind_result_v_o_top_bulk_data,omitempty"`
}

var poolTaobaoUniversalbpCrowdFindlistTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpCrowdFindlistTopResult)
	},
}

// GetTaobaoUniversalbpCrowdFindlistTopResult() 从对象池中获取TaobaoUniversalbpCrowdFindlistTopResult
func GetTaobaoUniversalbpCrowdFindlistTopResult() *TaobaoUniversalbpCrowdFindlistTopResult {
	return poolTaobaoUniversalbpCrowdFindlistTopResult.Get().(*TaobaoUniversalbpCrowdFindlistTopResult)
}

// ReleaseTaobaoUniversalbpCrowdFindlistTopResult 释放TaobaoUniversalbpCrowdFindlistTopResult
func ReleaseTaobaoUniversalbpCrowdFindlistTopResult(v *TaobaoUniversalbpCrowdFindlistTopResult) {
	v.Info = nil
	v.CrowdBindResultVOTopBulkData = nil
	poolTaobaoUniversalbpCrowdFindlistTopResult.Put(v)
}
