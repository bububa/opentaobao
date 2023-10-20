package ascpchannel

import (
	"sync"
)

// Futureinventorymainoperationquest 结构体
type Futureinventorymainoperationquest struct {
	// 子单操作明细列表
	DetailOperationDtos []Detailoperationdtos `json:"detail_operation_dtos,omitempty" xml:"detail_operation_dtos>detailoperationdtos,omitempty"`
	// 操作主单
	MainOrderDto *Mainorderdto `json:"main_order_dto,omitempty" xml:"main_order_dto,omitempty"`
}

var poolFutureinventorymainoperationquest = sync.Pool{
	New: func() any {
		return new(Futureinventorymainoperationquest)
	},
}

// GetFutureinventorymainoperationquest() 从对象池中获取Futureinventorymainoperationquest
func GetFutureinventorymainoperationquest() *Futureinventorymainoperationquest {
	return poolFutureinventorymainoperationquest.Get().(*Futureinventorymainoperationquest)
}

// ReleaseFutureinventorymainoperationquest 释放Futureinventorymainoperationquest
func ReleaseFutureinventorymainoperationquest(v *Futureinventorymainoperationquest) {
	v.DetailOperationDtos = v.DetailOperationDtos[:0]
	v.MainOrderDto = nil
	poolFutureinventorymainoperationquest.Put(v)
}
