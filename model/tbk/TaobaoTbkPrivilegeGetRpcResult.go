package tbk

import (
	"sync"
)

// TaobaoTbkPrivilegeGetRpcResult 结构体
type TaobaoTbkPrivilegeGetRpcResult struct {
	// data
	Data *MaterialDto `json:"data,omitempty" xml:"data,omitempty"`
}

var poolTaobaoTbkPrivilegeGetRpcResult = sync.Pool{
	New: func() any {
		return new(TaobaoTbkPrivilegeGetRpcResult)
	},
}

// GetTaobaoTbkPrivilegeGetRpcResult() 从对象池中获取TaobaoTbkPrivilegeGetRpcResult
func GetTaobaoTbkPrivilegeGetRpcResult() *TaobaoTbkPrivilegeGetRpcResult {
	return poolTaobaoTbkPrivilegeGetRpcResult.Get().(*TaobaoTbkPrivilegeGetRpcResult)
}

// ReleaseTaobaoTbkPrivilegeGetRpcResult 释放TaobaoTbkPrivilegeGetRpcResult
func ReleaseTaobaoTbkPrivilegeGetRpcResult(v *TaobaoTbkPrivilegeGetRpcResult) {
	v.Data = nil
	poolTaobaoTbkPrivilegeGetRpcResult.Put(v)
}
