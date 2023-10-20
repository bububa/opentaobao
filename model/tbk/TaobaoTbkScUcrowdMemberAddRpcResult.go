package tbk

import (
	"sync"
)

// TaobaoTbkScUcrowdMemberAddRpcResult 结构体
type TaobaoTbkScUcrowdMemberAddRpcResult struct {
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoTbkScUcrowdMemberAddRpcResult = sync.Pool{
	New: func() any {
		return new(TaobaoTbkScUcrowdMemberAddRpcResult)
	},
}

// GetTaobaoTbkScUcrowdMemberAddRpcResult() 从对象池中获取TaobaoTbkScUcrowdMemberAddRpcResult
func GetTaobaoTbkScUcrowdMemberAddRpcResult() *TaobaoTbkScUcrowdMemberAddRpcResult {
	return poolTaobaoTbkScUcrowdMemberAddRpcResult.Get().(*TaobaoTbkScUcrowdMemberAddRpcResult)
}

// ReleaseTaobaoTbkScUcrowdMemberAddRpcResult 释放TaobaoTbkScUcrowdMemberAddRpcResult
func ReleaseTaobaoTbkScUcrowdMemberAddRpcResult(v *TaobaoTbkScUcrowdMemberAddRpcResult) {
	v.Success = false
	poolTaobaoTbkScUcrowdMemberAddRpcResult.Put(v)
}
