package tbk

import (
	"sync"
)

// TaobaoTbkScUcrowdGetRpcResult 结构体
type TaobaoTbkScUcrowdGetRpcResult struct {
	// 人群
	Ucrowd *CrowdDto `json:"ucrowd,omitempty" xml:"ucrowd,omitempty"`
}

var poolTaobaoTbkScUcrowdGetRpcResult = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScUcrowdGetRpcResult)
	},
}

// GetTaobaoTbkScUcrowdGetRpcResult() 从对象池中获取TaobaoTbkScUcrowdGetRpcResult
func GetTaobaoTbkScUcrowdGetRpcResult() *TaobaoTbkScUcrowdGetRpcResult {
	return poolTaobaoTbkScUcrowdGetRpcResult.Get().(*TaobaoTbkScUcrowdGetRpcResult)
}

// ReleaseTaobaoTbkScUcrowdGetRpcResult 释放TaobaoTbkScUcrowdGetRpcResult
func ReleaseTaobaoTbkScUcrowdGetRpcResult(v *TaobaoTbkScUcrowdGetRpcResult) {
	v.Ucrowd = nil
	poolTaobaoTbkScUcrowdGetRpcResult.Put(v)
}
