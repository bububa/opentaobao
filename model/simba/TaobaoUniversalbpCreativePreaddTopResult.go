package simba

import (
	"sync"
)

// TaobaoUniversalbpCreativePreaddTopResult 结构体
type TaobaoUniversalbpCreativePreaddTopResult struct {
	// 请求系统信息
	Info *TopInfo `json:"info,omitempty" xml:"info,omitempty"`
	// 结果集
	PreAddItemCreativeVO *PreAddItemCreativeVo `json:"pre_add_item_creative_v_o,omitempty" xml:"pre_add_item_creative_v_o,omitempty"`
}

var poolTaobaoUniversalbpCreativePreaddTopResult = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpCreativePreaddTopResult)
	},
}

// GetTaobaoUniversalbpCreativePreaddTopResult() 从对象池中获取TaobaoUniversalbpCreativePreaddTopResult
func GetTaobaoUniversalbpCreativePreaddTopResult() *TaobaoUniversalbpCreativePreaddTopResult {
	return poolTaobaoUniversalbpCreativePreaddTopResult.Get().(*TaobaoUniversalbpCreativePreaddTopResult)
}

// ReleaseTaobaoUniversalbpCreativePreaddTopResult 释放TaobaoUniversalbpCreativePreaddTopResult
func ReleaseTaobaoUniversalbpCreativePreaddTopResult(v *TaobaoUniversalbpCreativePreaddTopResult) {
	v.Info = nil
	v.PreAddItemCreativeVO = nil
	poolTaobaoUniversalbpCreativePreaddTopResult.Put(v)
}
