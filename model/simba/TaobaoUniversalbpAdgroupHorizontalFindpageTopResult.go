package simba

import (
	"sync"
)

// TaobaoUniversalbpAdgroupHorizontalFindpageTopResult 结构体
type TaobaoUniversalbpAdgroupHorizontalFindpageTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	AdgroupVOTopBulkData *TopBulkData `json:"adgroup_v_o_top_bulk_data,omitempty" xml:"adgroup_v_o_top_bulk_data,omitempty"`
}

var poolTaobaoUniversalbpAdgroupHorizontalFindpageTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpAdgroupHorizontalFindpageTopResult)
	},
}

// GetTaobaoUniversalbpAdgroupHorizontalFindpageTopResult() 从对象池中获取TaobaoUniversalbpAdgroupHorizontalFindpageTopResult
func GetTaobaoUniversalbpAdgroupHorizontalFindpageTopResult() *TaobaoUniversalbpAdgroupHorizontalFindpageTopResult {
	return poolTaobaoUniversalbpAdgroupHorizontalFindpageTopResult.Get().(*TaobaoUniversalbpAdgroupHorizontalFindpageTopResult)
}

// ReleaseTaobaoUniversalbpAdgroupHorizontalFindpageTopResult 释放TaobaoUniversalbpAdgroupHorizontalFindpageTopResult
func ReleaseTaobaoUniversalbpAdgroupHorizontalFindpageTopResult(v *TaobaoUniversalbpAdgroupHorizontalFindpageTopResult) {
	v.Info = nil
	v.AdgroupVOTopBulkData = nil
	poolTaobaoUniversalbpAdgroupHorizontalFindpageTopResult.Put(v)
}
