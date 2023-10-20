package wdk

import (
	"sync"
)

// WarehouseOrderResponse 结构体
type WarehouseOrderResponse struct {
	// 子单列表
	SubOrders []WarehouseSubOrderResponse `json:"sub_orders,omitempty" xml:"sub_orders>warehouse_sub_order_response,omitempty"`
	// 商家编码
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 门店编码
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 渠道订单ID
	OutOrderId string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 仓作业单状态描述
	WarehouseStatusDesc string `json:"warehouse_status_desc,omitempty" xml:"warehouse_status_desc,omitempty"`
	// 仓单异常状态 (正常/缺货出/任务取消)
	ExcStatusDesc string `json:"exc_status_desc,omitempty" xml:"exc_status_desc,omitempty"`
	// 仓作业任务下发时间(仓接单)
	TaskDispatchedTime string `json:"task_dispatched_time,omitempty" xml:"task_dispatched_time,omitempty"`
	// 仓作业任务生成时间(准备开始作业)
	TaskGenerateTime string `json:"task_generate_time,omitempty" xml:"task_generate_time,omitempty"`
	// 拣货人员ID
	PickWorkerId string `json:"pick_worker_id,omitempty" xml:"pick_worker_id,omitempty"`
	// 拣货人员姓名
	PickWorkerName string `json:"pick_worker_name,omitempty" xml:"pick_worker_name,omitempty"`
	// 拣货开始时间(人员开始拣货)
	PickStartTime string `json:"pick_start_time,omitempty" xml:"pick_start_time,omitempty"`
	// 拣货结束时间
	PickFinishTime string `json:"pick_finish_time,omitempty" xml:"pick_finish_time,omitempty"`
	// 打包人员ID
	PackWorkerId string `json:"pack_worker_id,omitempty" xml:"pack_worker_id,omitempty"`
	// 打包人员姓名
	PackWorkerName string `json:"pack_worker_name,omitempty" xml:"pack_worker_name,omitempty"`
	// 打包开始时间
	PackStartTime string `json:"pack_start_time,omitempty" xml:"pack_start_time,omitempty"`
	// 打包结束时间
	PackFinishTime string `json:"pack_finish_time,omitempty" xml:"pack_finish_time,omitempty"`
	// 用户选择最晚送达时间
	LatestArrivalTime string `json:"latest_arrival_time,omitempty" xml:"latest_arrival_time,omitempty"`
	// 最晚出库时间
	LatestOutboundTime string `json:"latest_outbound_time,omitempty" xml:"latest_outbound_time,omitempty"`
	// 实际出库时间
	ActualOutboundTime string `json:"actual_outbound_time,omitempty" xml:"actual_outbound_time,omitempty"`
	// 仓作业取消时间
	CancelTime string `json:"cancel_time,omitempty" xml:"cancel_time,omitempty"`
	// 订单ID
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 订单来源 (2=美团/3=饿了么/26=京东到家/31=淘鲜达/28=私域渠道)
	OrderFrom int64 `json:"order_from,omitempty" xml:"order_from,omitempty"`
	// 仓作业单状态 �1 = 任务生成 2 = 拣货开始 3 = 拣货完成 4 = 打包开始 5 = 打包完成 -1 = 任务取消
	WarehouseStatus int64 `json:"warehouse_status,omitempty" xml:"warehouse_status,omitempty"`
	// 订单配送时效  1 = 自提 / 2 = 小时达
	DeliveryTimeMind int64 `json:"delivery_time_mind,omitempty" xml:"delivery_time_mind,omitempty"`
}

var poolWarehouseOrderResponse = sync.Pool{
	New: func() any {
		return new(WarehouseOrderResponse)
	},
}

// GetWarehouseOrderResponse() 从对象池中获取WarehouseOrderResponse
func GetWarehouseOrderResponse() *WarehouseOrderResponse {
	return poolWarehouseOrderResponse.Get().(*WarehouseOrderResponse)
}

// ReleaseWarehouseOrderResponse 释放WarehouseOrderResponse
func ReleaseWarehouseOrderResponse(v *WarehouseOrderResponse) {
	v.SubOrders = v.SubOrders[:0]
	v.MerchantCode = ""
	v.StoreId = ""
	v.OutOrderId = ""
	v.WarehouseStatusDesc = ""
	v.ExcStatusDesc = ""
	v.TaskDispatchedTime = ""
	v.TaskGenerateTime = ""
	v.PickWorkerId = ""
	v.PickWorkerName = ""
	v.PickStartTime = ""
	v.PickFinishTime = ""
	v.PackWorkerId = ""
	v.PackWorkerName = ""
	v.PackStartTime = ""
	v.PackFinishTime = ""
	v.LatestArrivalTime = ""
	v.LatestOutboundTime = ""
	v.ActualOutboundTime = ""
	v.CancelTime = ""
	v.BizOrderId = 0
	v.OrderFrom = 0
	v.WarehouseStatus = 0
	v.DeliveryTimeMind = 0
	poolWarehouseOrderResponse.Put(v)
}
