package cloudgame

import (
	"sync"
)

// AppDeletionCallbackRequest 结构体
type AppDeletionCallbackRequest struct {
	// 游戏删除结果
	Result *DeleteAppCommandResult `json:"result,omitempty" xml:"result,omitempty"`
}

var poolAppDeletionCallbackRequest = sync.Pool{
	New: func() any {
		return new(AppDeletionCallbackRequest)
	},
}

// GetAppDeletionCallbackRequest() 从对象池中获取AppDeletionCallbackRequest
func GetAppDeletionCallbackRequest() *AppDeletionCallbackRequest {
	return poolAppDeletionCallbackRequest.Get().(*AppDeletionCallbackRequest)
}

// ReleaseAppDeletionCallbackRequest 释放AppDeletionCallbackRequest
func ReleaseAppDeletionCallbackRequest(v *AppDeletionCallbackRequest) {
	v.Result = nil
	poolAppDeletionCallbackRequest.Put(v)
}
