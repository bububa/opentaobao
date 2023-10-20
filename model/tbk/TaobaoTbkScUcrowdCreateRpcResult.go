package tbk

import (
	"sync"
)

// TaobaoTbkScUcrowdCreateRpcResult 结构体
type TaobaoTbkScUcrowdCreateRpcResult struct {
	// 人群标签id
	UcrowdId int64 `json:"ucrowd_id,omitempty" xml:"ucrowd_id,omitempty"`
}

var poolTaobaoTbkScUcrowdCreateRpcResult = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScUcrowdCreateRpcResult)
	},
}

// GetTaobaoTbkScUcrowdCreateRpcResult() 从对象池中获取TaobaoTbkScUcrowdCreateRpcResult
func GetTaobaoTbkScUcrowdCreateRpcResult() *TaobaoTbkScUcrowdCreateRpcResult {
	return poolTaobaoTbkScUcrowdCreateRpcResult.Get().(*TaobaoTbkScUcrowdCreateRpcResult)
}

// ReleaseTaobaoTbkScUcrowdCreateRpcResult 释放TaobaoTbkScUcrowdCreateRpcResult
func ReleaseTaobaoTbkScUcrowdCreateRpcResult(v *TaobaoTbkScUcrowdCreateRpcResult) {
	v.UcrowdId = 0
	poolTaobaoTbkScUcrowdCreateRpcResult.Put(v)
}
