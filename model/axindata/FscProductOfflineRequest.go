package axindata

import (
	"sync"
)

// FscProductOfflineRequest 结构体
type FscProductOfflineRequest struct {
	// 产品编号列表
	ProductCodeList []string `json:"product_code_list,omitempty" xml:"product_code_list>string,omitempty"`
	// 供应商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
}

var poolFscProductOfflineRequest = sync.Pool{
	New: func() any {
		return new(FscProductOfflineRequest)
	},
}

// GetFscProductOfflineRequest() 从对象池中获取FscProductOfflineRequest
func GetFscProductOfflineRequest() *FscProductOfflineRequest {
	return poolFscProductOfflineRequest.Get().(*FscProductOfflineRequest)
}

// ReleaseFscProductOfflineRequest 释放FscProductOfflineRequest
func ReleaseFscProductOfflineRequest(v *FscProductOfflineRequest) {
	v.ProductCodeList = v.ProductCodeList[:0]
	v.SupplierId = ""
	poolFscProductOfflineRequest.Put(v)
}
