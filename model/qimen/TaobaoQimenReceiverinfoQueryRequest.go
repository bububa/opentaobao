package qimen

// TaobaoqimenreceiverinfoqueryRequest 结构体
type TaobaoqimenreceiverinfoqueryRequest struct {
	// 订单收件人 ID, string (50)
	Oaid string `json:"oaid,omitempty" xml:"oaid,omitempty"`
	// 出库单号, string (50) , 必填
	DeliveryOrderCode string `json:"deliveryOrderCode,omitempty" xml:"deliveryOrderCode,omitempty"`
	// 货主ID
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 仓库编码
	WarehouseCode string `json:"warehouseCode,omitempty" xml:"warehouseCode,omitempty"`
	// 使用场景。1001，顺丰电子面单发货；1002，4通一达电子面单发货；1003，EMS电子面单发货；1004，其他电子面单发货；2001，客户售后服务
	Scene string `json:"scene,omitempty" xml:"scene,omitempty"`
}
