package wlbimports

import (
	"sync"
)

// TranStoreResult 结构体
type TranStoreResult struct {
	// 中转仓代码
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// 中转仓名字
	StoreName string `json:"store_name,omitempty" xml:"store_name,omitempty"`
	// 中转仓地址
	StoreAddress string `json:"store_address,omitempty" xml:"store_address,omitempty"`
}

var poolTranStoreResult = sync.Pool{
	New: func() any {
		return new(TranStoreResult)
	},
}

// GetTranStoreResult() 从对象池中获取TranStoreResult
func GetTranStoreResult() *TranStoreResult {
	return poolTranStoreResult.Get().(*TranStoreResult)
}

// ReleaseTranStoreResult 释放TranStoreResult
func ReleaseTranStoreResult(v *TranStoreResult) {
	v.StoreCode = ""
	v.StoreName = ""
	v.StoreAddress = ""
	poolTranStoreResult.Put(v)
}
