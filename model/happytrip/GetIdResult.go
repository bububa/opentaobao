package happytrip

import (
	"sync"
)

// GetIdResult 结构体
type GetIdResult struct {
	// 返回的id
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

var poolGetIdResult = sync.Pool{
	New: func() any {
		return new(GetIdResult)
	},
}

// GetGetIdResult() 从对象池中获取GetIdResult
func GetGetIdResult() *GetIdResult {
	return poolGetIdResult.Get().(*GetIdResult)
}

// ReleaseGetIdResult 释放GetIdResult
func ReleaseGetIdResult(v *GetIdResult) {
	v.OrderId = ""
	poolGetIdResult.Put(v)
}
