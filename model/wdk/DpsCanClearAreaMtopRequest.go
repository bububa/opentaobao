package wdk

import (
	"sync"
)

// DpsCanClearAreaMtopRequest 结构体
type DpsCanClearAreaMtopRequest struct {
	// 波次号
	WaveCode string `json:"wave_code,omitempty" xml:"wave_code,omitempty"`
	// 仓code
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
}

var poolDpsCanClearAreaMtopRequest = sync.Pool{
	New: func() any {
		return new(DpsCanClearAreaMtopRequest)
	},
}

// GetDpsCanClearAreaMtopRequest() 从对象池中获取DpsCanClearAreaMtopRequest
func GetDpsCanClearAreaMtopRequest() *DpsCanClearAreaMtopRequest {
	return poolDpsCanClearAreaMtopRequest.Get().(*DpsCanClearAreaMtopRequest)
}

// ReleaseDpsCanClearAreaMtopRequest 释放DpsCanClearAreaMtopRequest
func ReleaseDpsCanClearAreaMtopRequest(v *DpsCanClearAreaMtopRequest) {
	v.WaveCode = ""
	v.WarehouseCode = ""
	poolDpsCanClearAreaMtopRequest.Put(v)
}
