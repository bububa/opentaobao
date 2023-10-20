package perfect

import (
	"sync"
)

// LoadReceiveRequest 结构体
type LoadReceiveRequest struct {
	// 装箱单
	ContainerOrders []LoadContainerOrderRequest `json:"container_orders,omitempty" xml:"container_orders>load_container_order_request,omitempty"`
	// 操作时间
	OperateTime string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
	// 扩展
	Attributes string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 操作人
	OperatorCode string `json:"operator_code,omitempty" xml:"operator_code,omitempty"`
	// 仓code
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
}

var poolLoadReceiveRequest = sync.Pool{
	New: func() any {
		return new(LoadReceiveRequest)
	},
}

// GetLoadReceiveRequest() 从对象池中获取LoadReceiveRequest
func GetLoadReceiveRequest() *LoadReceiveRequest {
	return poolLoadReceiveRequest.Get().(*LoadReceiveRequest)
}

// ReleaseLoadReceiveRequest 释放LoadReceiveRequest
func ReleaseLoadReceiveRequest(v *LoadReceiveRequest) {
	v.ContainerOrders = v.ContainerOrders[:0]
	v.OperateTime = ""
	v.Attributes = ""
	v.OperatorCode = ""
	v.WarehouseCode = ""
	poolLoadReceiveRequest.Put(v)
}
