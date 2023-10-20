package idle

import (
	"sync"
)

// TopListResult 结构体
type TopListResult struct {
	// 移除失败商品信息列表
	Data []IdleItemAutoRechargeBatchRemoveResult `json:"data,omitempty" xml:"data>idle_item_auto_recharge_batch_remove_result,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误说明
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 请求是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTopListResult = sync.Pool{
	New: func() any {
		return new(TopListResult)
	},
}

// GetTopListResult() 从对象池中获取TopListResult
func GetTopListResult() *TopListResult {
	return poolTopListResult.Get().(*TopListResult)
}

// ReleaseTopListResult 释放TopListResult
func ReleaseTopListResult(v *TopListResult) {
	v.Data = v.Data[:0]
	v.ErrCode = ""
	v.ErrMsg = ""
	v.Success = false
	poolTopListResult.Put(v)
}
