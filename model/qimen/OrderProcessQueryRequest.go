package qimen

import (
	"sync"
)

// OrderProcessQueryRequest 结构体
type OrderProcessQueryRequest struct {
	// 单据类型(JYCK=一般交易出库单;HHCK=换货出库;BFCK=补发出库;PTCK=普通出库单;DBCK=调拨出库;QTCK=其他出库;B2BRK=B2B入库;B2BCK=B2B出库;CGRK=采购入库;DBRK=调拨入库;QTRK=其他入库;XTRK=销退入库;HHRK=换货入库;CNJG=仓内加工单)
	OrderType string `json:"orderType,omitempty" xml:"orderType,omitempty"`
	// 单据号
	OrderCode string `json:"orderCode,omitempty" xml:"orderCode,omitempty"`
	// 仓储系统单据号
	OrderId string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// 仓库编码
	WarehouseCode string `json:"warehouseCode,omitempty" xml:"warehouseCode,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 交易单号
	OrderSourceCode string `json:"orderSourceCode,omitempty" xml:"orderSourceCode,omitempty"`
	// 货主编码
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenOrderprocessQueryMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}

var poolOrderProcessQueryRequest = sync.Pool{
	New: func() any {
		return new(OrderProcessQueryRequest)
	},
}

// GetOrderProcessQueryRequest() 从对象池中获取OrderProcessQueryRequest
func GetOrderProcessQueryRequest() *OrderProcessQueryRequest {
	return poolOrderProcessQueryRequest.Get().(*OrderProcessQueryRequest)
}

// ReleaseOrderProcessQueryRequest 释放OrderProcessQueryRequest
func ReleaseOrderProcessQueryRequest(v *OrderProcessQueryRequest) {
	v.OrderType = ""
	v.OrderCode = ""
	v.OrderId = ""
	v.WarehouseCode = ""
	v.Remark = ""
	v.OrderSourceCode = ""
	v.OwnerCode = ""
	v.ExtendProps = nil
	poolOrderProcessQueryRequest.Put(v)
}
