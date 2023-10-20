package idle

import (
	"sync"
)

// IdleItemAutoRechargeBatchRemoveResult 结构体
type IdleItemAutoRechargeBatchRemoveResult struct {
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误说明
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 移除失败商品id
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

var poolIdleItemAutoRechargeBatchRemoveResult = sync.Pool{
	New: func() any {
		return new(IdleItemAutoRechargeBatchRemoveResult)
	},
}

// GetIdleItemAutoRechargeBatchRemoveResult() 从对象池中获取IdleItemAutoRechargeBatchRemoveResult
func GetIdleItemAutoRechargeBatchRemoveResult() *IdleItemAutoRechargeBatchRemoveResult {
	return poolIdleItemAutoRechargeBatchRemoveResult.Get().(*IdleItemAutoRechargeBatchRemoveResult)
}

// ReleaseIdleItemAutoRechargeBatchRemoveResult 释放IdleItemAutoRechargeBatchRemoveResult
func ReleaseIdleItemAutoRechargeBatchRemoveResult(v *IdleItemAutoRechargeBatchRemoveResult) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.ItemId = 0
	poolIdleItemAutoRechargeBatchRemoveResult.Put(v)
}
