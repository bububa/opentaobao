package koubeimall

import (
	"sync"
)

// TaobaoKoubeiMallCommonStoreCommentPageResult 结构体
type TaobaoKoubeiMallCommonStoreCommentPageResult struct {
	// API请求全链路追踪ID
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// API返回的分页模型
	Data *PageResult `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息
	Error *TribeError `json:"error,omitempty" xml:"error,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoKoubeiMallCommonStoreCommentPageResult = sync.Pool{
	New: func() any {
		return new(TaobaoKoubeiMallCommonStoreCommentPageResult)
	},
}

// GetTaobaoKoubeiMallCommonStoreCommentPageResult() 从对象池中获取TaobaoKoubeiMallCommonStoreCommentPageResult
func GetTaobaoKoubeiMallCommonStoreCommentPageResult() *TaobaoKoubeiMallCommonStoreCommentPageResult {
	return poolTaobaoKoubeiMallCommonStoreCommentPageResult.Get().(*TaobaoKoubeiMallCommonStoreCommentPageResult)
}

// ReleaseTaobaoKoubeiMallCommonStoreCommentPageResult 释放TaobaoKoubeiMallCommonStoreCommentPageResult
func ReleaseTaobaoKoubeiMallCommonStoreCommentPageResult(v *TaobaoKoubeiMallCommonStoreCommentPageResult) {
	v.TraceId = ""
	v.Data = nil
	v.Error = nil
	v.Success = false
	poolTaobaoKoubeiMallCommonStoreCommentPageResult.Put(v)
}
