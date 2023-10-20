package koubeimall

import (
	"sync"
)

// TaobaoKoubeiMallCommonMallNearListResult 结构体
type TaobaoKoubeiMallCommonMallNearListResult struct {
	// 附近商圈列表模型
	MallList []MallDto `json:"mall_list,omitempty" xml:"mall_list>mall_dto,omitempty"`
	// API请求全链路追踪ID
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 错误信息，success=false时，返回相关错误信息
	Error *TribeError `json:"error,omitempty" xml:"error,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoKoubeiMallCommonMallNearListResult = sync.Pool{
	New: func() any {
		return new(TaobaoKoubeiMallCommonMallNearListResult)
	},
}

// GetTaobaoKoubeiMallCommonMallNearListResult() 从对象池中获取TaobaoKoubeiMallCommonMallNearListResult
func GetTaobaoKoubeiMallCommonMallNearListResult() *TaobaoKoubeiMallCommonMallNearListResult {
	return poolTaobaoKoubeiMallCommonMallNearListResult.Get().(*TaobaoKoubeiMallCommonMallNearListResult)
}

// ReleaseTaobaoKoubeiMallCommonMallNearListResult 释放TaobaoKoubeiMallCommonMallNearListResult
func ReleaseTaobaoKoubeiMallCommonMallNearListResult(v *TaobaoKoubeiMallCommonMallNearListResult) {
	v.MallList = v.MallList[:0]
	v.TraceId = ""
	v.Error = nil
	v.Success = false
	poolTaobaoKoubeiMallCommonMallNearListResult.Put(v)
}
