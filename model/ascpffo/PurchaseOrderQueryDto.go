package ascpffo

import (
	"sync"
)

// PurchaseOrderQueryDto 结构体
type PurchaseOrderQueryDto struct {
	// 采购单号
	PurchaseOrderNo string `json:"purchase_order_no,omitempty" xml:"purchase_order_no,omitempty"`
	// 行业账套编码
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
}

var poolPurchaseOrderQueryDto = sync.Pool{
	New: func() any {
		return new(PurchaseOrderQueryDto)
	},
}

// GetPurchaseOrderQueryDto() 从对象池中获取PurchaseOrderQueryDto
func GetPurchaseOrderQueryDto() *PurchaseOrderQueryDto {
	return poolPurchaseOrderQueryDto.Get().(*PurchaseOrderQueryDto)
}

// ReleasePurchaseOrderQueryDto 释放PurchaseOrderQueryDto
func ReleasePurchaseOrderQueryDto(v *PurchaseOrderQueryDto) {
	v.PurchaseOrderNo = ""
	v.BizType = 0
	poolPurchaseOrderQueryDto.Put(v)
}
