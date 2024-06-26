package qimen

import (
	"sync"
)

// OrderPendingRequest 结构体
type OrderPendingRequest struct {
	// 操作类型(pending=挂起;restore=恢复)
	ActionType string `json:"actionType,omitempty" xml:"actionType,omitempty"`
	// 仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER)
	WarehouseCode string `json:"warehouseCode,omitempty" xml:"warehouseCode,omitempty"`
	// 货主编码
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 单据编码
	OrderCode string `json:"orderCode,omitempty" xml:"orderCode,omitempty"`
	// 仓储系统单据编码
	OrderId string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// 单据类型(JYCK=一般交易出库单;HHCK=换货出库;BFCK=补发出库;PTCK=普通出库单;DBCK=调拨出库;QTCK=其他出库;B2BRK=B2B入库;B2BCK=B2B出库;CGRK=采购入库;DBRK=调拨入库;QTRK=其他入库;XTRK=销退入库;HHRK=换货入库;CNJG=仓内加工单)
	OrderType string `json:"orderType,omitempty" xml:"orderType,omitempty"`
	// 挂起/恢复原因
	Reason string `json:"reason,omitempty" xml:"reason,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenOrderPendingMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}

var poolOrderPendingRequest = sync.Pool{
	New: func() any {
		return new(OrderPendingRequest)
	},
}

// GetOrderPendingRequest() 从对象池中获取OrderPendingRequest
func GetOrderPendingRequest() *OrderPendingRequest {
	return poolOrderPendingRequest.Get().(*OrderPendingRequest)
}

// ReleaseOrderPendingRequest 释放OrderPendingRequest
func ReleaseOrderPendingRequest(v *OrderPendingRequest) {
	v.ActionType = ""
	v.WarehouseCode = ""
	v.OwnerCode = ""
	v.OrderCode = ""
	v.OrderId = ""
	v.OrderType = ""
	v.Reason = ""
	v.ExtendProps = nil
	poolOrderPendingRequest.Put(v)
}
