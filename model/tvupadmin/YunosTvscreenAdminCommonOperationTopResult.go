package tvupadmin

import (
	"sync"
)

// YunosTvscreenAdminCommonOperationTopResult 结构体
type YunosTvscreenAdminCommonOperationTopResult struct {
	// result
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolYunosTvscreenAdminCommonOperationTopResult = sync.Pool{
	New: func() any {
		return new(YunosTvscreenAdminCommonOperationTopResult)
	},
}

// GetYunosTvscreenAdminCommonOperationTopResult() 从对象池中获取YunosTvscreenAdminCommonOperationTopResult
func GetYunosTvscreenAdminCommonOperationTopResult() *YunosTvscreenAdminCommonOperationTopResult {
	return poolYunosTvscreenAdminCommonOperationTopResult.Get().(*YunosTvscreenAdminCommonOperationTopResult)
}

// ReleaseYunosTvscreenAdminCommonOperationTopResult 释放YunosTvscreenAdminCommonOperationTopResult
func ReleaseYunosTvscreenAdminCommonOperationTopResult(v *YunosTvscreenAdminCommonOperationTopResult) {
	v.Result = ""
	v.Message = ""
	poolYunosTvscreenAdminCommonOperationTopResult.Put(v)
}
