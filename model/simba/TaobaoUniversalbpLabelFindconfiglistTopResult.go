package simba

import (
	"sync"
)

// TaobaoUniversalbpLabelFindconfiglistTopResult 结构体
type TaobaoUniversalbpLabelFindconfiglistTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	LabelConfigVOTopBulkData *TopBulkData `json:"label_config_v_o_top_bulk_data,omitempty" xml:"label_config_v_o_top_bulk_data,omitempty"`
}

var poolTaobaoUniversalbpLabelFindconfiglistTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpLabelFindconfiglistTopResult)
	},
}

// GetTaobaoUniversalbpLabelFindconfiglistTopResult() 从对象池中获取TaobaoUniversalbpLabelFindconfiglistTopResult
func GetTaobaoUniversalbpLabelFindconfiglistTopResult() *TaobaoUniversalbpLabelFindconfiglistTopResult {
	return poolTaobaoUniversalbpLabelFindconfiglistTopResult.Get().(*TaobaoUniversalbpLabelFindconfiglistTopResult)
}

// ReleaseTaobaoUniversalbpLabelFindconfiglistTopResult 释放TaobaoUniversalbpLabelFindconfiglistTopResult
func ReleaseTaobaoUniversalbpLabelFindconfiglistTopResult(v *TaobaoUniversalbpLabelFindconfiglistTopResult) {
	v.Info = nil
	v.LabelConfigVOTopBulkData = nil
	poolTaobaoUniversalbpLabelFindconfiglistTopResult.Put(v)
}
