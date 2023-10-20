package wdk

import (
	"sync"
)

// ErpBillCallbackDto 结构体
type ErpBillCallbackDto struct {
	// s失败原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 0：失败，1：成功
	Success string `json:"success,omitempty" xml:"success,omitempty"`
	// 7：退货单   3：调拨出库单, 10 质量反馈
	BillType string `json:"bill_type,omitempty" xml:"bill_type,omitempty"`
	// 单据号
	BizOrderCode string `json:"biz_order_code,omitempty" xml:"biz_order_code,omitempty"`
	// 唯一识别码
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// warehouseCode
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
}

var poolErpBillCallbackDto = sync.Pool{
	New: func() any {
		return new(ErpBillCallbackDto)
	},
}

// GetErpBillCallbackDto() 从对象池中获取ErpBillCallbackDto
func GetErpBillCallbackDto() *ErpBillCallbackDto {
	return poolErpBillCallbackDto.Get().(*ErpBillCallbackDto)
}

// ReleaseErpBillCallbackDto 释放ErpBillCallbackDto
func ReleaseErpBillCallbackDto(v *ErpBillCallbackDto) {
	v.Reason = ""
	v.Success = ""
	v.BillType = ""
	v.BizOrderCode = ""
	v.Uuid = ""
	v.WarehouseCode = ""
	poolErpBillCallbackDto.Put(v)
}
