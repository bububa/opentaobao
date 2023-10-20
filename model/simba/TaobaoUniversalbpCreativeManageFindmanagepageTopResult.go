package simba

import (
	"sync"
)

// TaobaoUniversalbpCreativeManageFindmanagepageTopResult 结构体
type TaobaoUniversalbpCreativeManageFindmanagepageTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	CreativeVOTopBulkData *TopBulkData `json:"creative_v_o_top_bulk_data,omitempty" xml:"creative_v_o_top_bulk_data,omitempty"`
}

var poolTaobaoUniversalbpCreativeManageFindmanagepageTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpCreativeManageFindmanagepageTopResult)
	},
}

// GetTaobaoUniversalbpCreativeManageFindmanagepageTopResult() 从对象池中获取TaobaoUniversalbpCreativeManageFindmanagepageTopResult
func GetTaobaoUniversalbpCreativeManageFindmanagepageTopResult() *TaobaoUniversalbpCreativeManageFindmanagepageTopResult {
	return poolTaobaoUniversalbpCreativeManageFindmanagepageTopResult.Get().(*TaobaoUniversalbpCreativeManageFindmanagepageTopResult)
}

// ReleaseTaobaoUniversalbpCreativeManageFindmanagepageTopResult 释放TaobaoUniversalbpCreativeManageFindmanagepageTopResult
func ReleaseTaobaoUniversalbpCreativeManageFindmanagepageTopResult(v *TaobaoUniversalbpCreativeManageFindmanagepageTopResult) {
	v.Info = nil
	v.CreativeVOTopBulkData = nil
	poolTaobaoUniversalbpCreativeManageFindmanagepageTopResult.Put(v)
}
