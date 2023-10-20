package simba

import (
	"sync"
)

// TaobaoUniversalbpMemberFindbrandinfolistTopResult 结构体
type TaobaoUniversalbpMemberFindbrandinfolistTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	BrandInfoVOTopBulkData *TopBulkData `json:"brand_info_v_o_top_bulk_data,omitempty" xml:"brand_info_v_o_top_bulk_data,omitempty"`
}

var poolTaobaoUniversalbpMemberFindbrandinfolistTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpMemberFindbrandinfolistTopResult)
	},
}

// GetTaobaoUniversalbpMemberFindbrandinfolistTopResult() 从对象池中获取TaobaoUniversalbpMemberFindbrandinfolistTopResult
func GetTaobaoUniversalbpMemberFindbrandinfolistTopResult() *TaobaoUniversalbpMemberFindbrandinfolistTopResult {
	return poolTaobaoUniversalbpMemberFindbrandinfolistTopResult.Get().(*TaobaoUniversalbpMemberFindbrandinfolistTopResult)
}

// ReleaseTaobaoUniversalbpMemberFindbrandinfolistTopResult 释放TaobaoUniversalbpMemberFindbrandinfolistTopResult
func ReleaseTaobaoUniversalbpMemberFindbrandinfolistTopResult(v *TaobaoUniversalbpMemberFindbrandinfolistTopResult) {
	v.Info = nil
	v.BrandInfoVOTopBulkData = nil
	poolTaobaoUniversalbpMemberFindbrandinfolistTopResult.Put(v)
}
