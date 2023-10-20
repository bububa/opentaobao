package qimen

import (
	"sync"
)

// DeliveryOrderQueryRequest 结构体
type DeliveryOrderQueryRequest struct {
	// 货主编码
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 仓库编码
	WarehouseCode string `json:"warehouseCode,omitempty" xml:"warehouseCode,omitempty"`
	// 发库单号
	OrderCode string `json:"orderCode,omitempty" xml:"orderCode,omitempty"`
	// 仓储系统发库单号
	OrderId string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// 交易单号
	OrderSourceCode string `json:"orderSourceCode,omitempty" xml:"orderSourceCode,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 当前页
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页orderLine条数(最多100条)
	PageSize int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenDeliveryorderQueryMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}

var poolDeliveryOrderQueryRequest = sync.Pool{
	New: func() any {
		return new(DeliveryOrderQueryRequest)
	},
}

// GetDeliveryOrderQueryRequest() 从对象池中获取DeliveryOrderQueryRequest
func GetDeliveryOrderQueryRequest() *DeliveryOrderQueryRequest {
	return poolDeliveryOrderQueryRequest.Get().(*DeliveryOrderQueryRequest)
}

// ReleaseDeliveryOrderQueryRequest 释放DeliveryOrderQueryRequest
func ReleaseDeliveryOrderQueryRequest(v *DeliveryOrderQueryRequest) {
	v.OwnerCode = ""
	v.WarehouseCode = ""
	v.OrderCode = ""
	v.OrderId = ""
	v.OrderSourceCode = ""
	v.Remark = ""
	v.Page = 0
	v.PageSize = 0
	v.ExtendProps = nil
	poolDeliveryOrderQueryRequest.Put(v)
}
