package tbk

import (
	"sync"
)

// TaobaoTbkScUcrowdDeleteRpcResult 结构体
type TaobaoTbkScUcrowdDeleteRpcResult struct {
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoTbkScUcrowdDeleteRpcResult = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScUcrowdDeleteRpcResult)
	},
}

// GetTaobaoTbkScUcrowdDeleteRpcResult() 从对象池中获取TaobaoTbkScUcrowdDeleteRpcResult
func GetTaobaoTbkScUcrowdDeleteRpcResult() *TaobaoTbkScUcrowdDeleteRpcResult {
	return poolTaobaoTbkScUcrowdDeleteRpcResult.Get().(*TaobaoTbkScUcrowdDeleteRpcResult)
}

// ReleaseTaobaoTbkScUcrowdDeleteRpcResult 释放TaobaoTbkScUcrowdDeleteRpcResult
func ReleaseTaobaoTbkScUcrowdDeleteRpcResult(v *TaobaoTbkScUcrowdDeleteRpcResult) {
	v.Success = false
	poolTaobaoTbkScUcrowdDeleteRpcResult.Put(v)
}
