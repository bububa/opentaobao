package qimen

import (
	"sync"
)

// OrderProcess 结构体
type OrderProcess struct {
	// 处理流程
	Processes []Process `json:"processes,omitempty" xml:"processes>process,omitempty"`
	// 单据号
	OrderCode string `json:"orderCode,omitempty" xml:"orderCode,omitempty"`
	// 仓储系统单据号
	OrderId string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// 单据类型(JYCK=一般交易出库单;HHCK=换货出库;BFCK=补发出库;PTCK=普通出库单;DBCK=调拨出库;QTCK=其他出库;B2BRK=B2B入库;B2BCK=B2B出库;CGRK=采购入库;DBRK=调拨入库;QTRK=其他入库;XTRK=销退入库;HHRK=换货入库;CNJG=仓内加工单)
	OrderType string `json:"orderType,omitempty" xml:"orderType,omitempty"`
	// 仓库编码
	WarehouseCode string `json:"warehouseCode,omitempty" xml:"warehouseCode,omitempty"`
}

var poolOrderProcess = sync.Pool{
	New: func() any {
		return new(OrderProcess)
	},
}

// GetOrderProcess() 从对象池中获取OrderProcess
func GetOrderProcess() *OrderProcess {
	return poolOrderProcess.Get().(*OrderProcess)
}

// ReleaseOrderProcess 释放OrderProcess
func ReleaseOrderProcess(v *OrderProcess) {
	v.Processes = v.Processes[:0]
	v.OrderCode = ""
	v.OrderId = ""
	v.OrderType = ""
	v.WarehouseCode = ""
	poolOrderProcess.Put(v)
}
