package axindata

import (
	"sync"
)

// FscProductOnlineRequest 结构体
type FscProductOnlineRequest struct {
	// 产品编号列表
	ProductCodeList []string `json:"product_code_list,omitempty" xml:"product_code_list>string,omitempty"`
	// 供应商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
}

var poolFscProductOnlineRequest = sync.Pool{
	New: func() any {
		return new(FscProductOnlineRequest)
	},
}

// GetFscProductOnlineRequest() 从对象池中获取FscProductOnlineRequest
func GetFscProductOnlineRequest() *FscProductOnlineRequest {
	return poolFscProductOnlineRequest.Get().(*FscProductOnlineRequest)
}

// ReleaseFscProductOnlineRequest 释放FscProductOnlineRequest
func ReleaseFscProductOnlineRequest(v *FscProductOnlineRequest) {
	v.ProductCodeList = v.ProductCodeList[:0]
	v.SupplierId = ""
	poolFscProductOnlineRequest.Put(v)
}
