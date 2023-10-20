package ascpffo

import (
	"sync"
)

// PurchaseOrderItemQueryDto 结构体
type PurchaseOrderItemQueryDto struct {
	// 采购单号
	PurchaseOrderNo string `json:"purchase_order_no,omitempty" xml:"purchase_order_no,omitempty"`
	// 账套编码
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 分页页码
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolPurchaseOrderItemQueryDto = sync.Pool{
	New: func() any {
		return new(PurchaseOrderItemQueryDto)
	},
}

// GetPurchaseOrderItemQueryDto() 从对象池中获取PurchaseOrderItemQueryDto
func GetPurchaseOrderItemQueryDto() *PurchaseOrderItemQueryDto {
	return poolPurchaseOrderItemQueryDto.Get().(*PurchaseOrderItemQueryDto)
}

// ReleasePurchaseOrderItemQueryDto 释放PurchaseOrderItemQueryDto
func ReleasePurchaseOrderItemQueryDto(v *PurchaseOrderItemQueryDto) {
	v.PurchaseOrderNo = ""
	v.BizType = 0
	v.PageIndex = 0
	v.PageSize = 0
	poolPurchaseOrderItemQueryDto.Put(v)
}
