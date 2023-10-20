package qimen

import (
	"sync"
)

// OrderInfo 结构体
type OrderInfo struct {
	// 订单创建时间(YYYY-MM-DD HH:MM:SS)
	CreateTime string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 出库单仓储系统编码
	DeliveryOrderId string `json:"deliveryOrderId,omitempty" xml:"deliveryOrderId,omitempty"`
	// 仓库编码(统仓统配使用)
	WarehouseCode string `json:"warehouseCode,omitempty" xml:"warehouseCode,omitempty"`
	// 物流公司编码(统仓统配使用)
	LogisticsCode string `json:"logisticsCode,omitempty" xml:"logisticsCode,omitempty"`
}

var poolOrderInfo = sync.Pool{
	New: func() any {
		return new(OrderInfo)
	},
}

// GetOrderInfo() 从对象池中获取OrderInfo
func GetOrderInfo() *OrderInfo {
	return poolOrderInfo.Get().(*OrderInfo)
}

// ReleaseOrderInfo 释放OrderInfo
func ReleaseOrderInfo(v *OrderInfo) {
	v.CreateTime = ""
	v.DeliveryOrderId = ""
	v.WarehouseCode = ""
	v.LogisticsCode = ""
	poolOrderInfo.Put(v)
}
