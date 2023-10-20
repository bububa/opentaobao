package qimen

import (
	"sync"
)

// EntryOrderQueryRequest 结构体
type EntryOrderQueryRequest struct {
	// 货主编码
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 仓库编码
	WarehouseCode string `json:"warehouseCode,omitempty" xml:"warehouseCode,omitempty"`
	// 入库单编码
	EntryOrderCode string `json:"entryOrderCode,omitempty" xml:"entryOrderCode,omitempty"`
	// 仓储系统入库单ID
	EntryOrderId string `json:"entryOrderId,omitempty" xml:"entryOrderId,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 当前页(从1开始)
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页orderLine条数(最多100条)
	PageSize int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenEntryorderQueryMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}

var poolEntryOrderQueryRequest = sync.Pool{
	New: func() any {
		return new(EntryOrderQueryRequest)
	},
}

// GetEntryOrderQueryRequest() 从对象池中获取EntryOrderQueryRequest
func GetEntryOrderQueryRequest() *EntryOrderQueryRequest {
	return poolEntryOrderQueryRequest.Get().(*EntryOrderQueryRequest)
}

// ReleaseEntryOrderQueryRequest 释放EntryOrderQueryRequest
func ReleaseEntryOrderQueryRequest(v *EntryOrderQueryRequest) {
	v.OwnerCode = ""
	v.WarehouseCode = ""
	v.EntryOrderCode = ""
	v.EntryOrderId = ""
	v.Remark = ""
	v.Page = 0
	v.PageSize = 0
	v.ExtendProps = nil
	poolEntryOrderQueryRequest.Put(v)
}
