package tvupadmin

import (
	"sync"
)

// YunosTradeAdminCommonOperationTopResult 结构体
type YunosTradeAdminCommonOperationTopResult struct {
	// 实际内容
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

var poolYunosTradeAdminCommonOperationTopResult = sync.Pool{
	New: func() any {
		return new(YunosTradeAdminCommonOperationTopResult)
	},
}

// GetYunosTradeAdminCommonOperationTopResult() 从对象池中获取YunosTradeAdminCommonOperationTopResult
func GetYunosTradeAdminCommonOperationTopResult() *YunosTradeAdminCommonOperationTopResult {
	return poolYunosTradeAdminCommonOperationTopResult.Get().(*YunosTradeAdminCommonOperationTopResult)
}

// ReleaseYunosTradeAdminCommonOperationTopResult 释放YunosTradeAdminCommonOperationTopResult
func ReleaseYunosTradeAdminCommonOperationTopResult(v *YunosTradeAdminCommonOperationTopResult) {
	v.Result = ""
	poolYunosTradeAdminCommonOperationTopResult.Put(v)
}
