package koubeimall

import (
	"sync"
)

// TaobaoKoubeiMallCommonStoreDisplayGoodsListResult 结构体
type TaobaoKoubeiMallCommonStoreDisplayGoodsListResult struct {
	// API请求全链路追踪ID
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// API返回的推荐菜data模型
	Data *DisplayGoodsDto `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息，success=false时，返回相关错误信息
	Error *TribeError `json:"error,omitempty" xml:"error,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoKoubeiMallCommonStoreDisplayGoodsListResult = sync.Pool{
	New: func() any {
		return new(TaobaoKoubeiMallCommonStoreDisplayGoodsListResult)
	},
}

// GetTaobaoKoubeiMallCommonStoreDisplayGoodsListResult() 从对象池中获取TaobaoKoubeiMallCommonStoreDisplayGoodsListResult
func GetTaobaoKoubeiMallCommonStoreDisplayGoodsListResult() *TaobaoKoubeiMallCommonStoreDisplayGoodsListResult {
	return poolTaobaoKoubeiMallCommonStoreDisplayGoodsListResult.Get().(*TaobaoKoubeiMallCommonStoreDisplayGoodsListResult)
}

// ReleaseTaobaoKoubeiMallCommonStoreDisplayGoodsListResult 释放TaobaoKoubeiMallCommonStoreDisplayGoodsListResult
func ReleaseTaobaoKoubeiMallCommonStoreDisplayGoodsListResult(v *TaobaoKoubeiMallCommonStoreDisplayGoodsListResult) {
	v.TraceId = ""
	v.Data = nil
	v.Error = nil
	v.Success = false
	poolTaobaoKoubeiMallCommonStoreDisplayGoodsListResult.Put(v)
}
